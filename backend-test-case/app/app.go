package app

import (
	"backend_test_case/config"
	"backend_test_case/model/dto"
	"backend_test_case/router"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func initEnv() (dto.ConfigData, error) {
	var configData dto.ConfigData

	if err := godotenv.Load(); err != nil {
		return dto.ConfigData{}, err
	}

	if port := os.Getenv("PORT"); port != "" {
		configData.AppConfig.Port = port
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbMaxIdle := os.Getenv("MAX_IDLE")
	dbMaxConn := os.Getenv("MAX_CONN")
	dbMaxLifeTime := os.Getenv("MAX_LIFE_TIME")
	dbLogMode := os.Getenv("LOG_MODE")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPass == "" || dbName == "" || dbMaxIdle == "" || dbMaxConn == "" || dbMaxLifeTime == "" || dbLogMode == "" {
		return dto.ConfigData{}, errors.New("DB config is not set")
	}

	maxIdle, err := strconv.Atoi(dbMaxIdle)
	if err != nil {
		return dto.ConfigData{}, err
	}

	maxConn, err := strconv.Atoi(dbMaxIdle)
	if err != nil {
		return dto.ConfigData{}, err
	}

	logMode, err := strconv.Atoi(dbLogMode)
	if err != nil {
		return dto.ConfigData{}, err
	}

	configData.DbConfig.Host = dbHost
	configData.DbConfig.Port = dbPort
	configData.DbConfig.User = dbUser
	configData.DbConfig.Pass = dbPass
	configData.DbConfig.Database = dbName
	configData.DbConfig.MaxIdle = maxIdle
	configData.DbConfig.MaxConn = maxConn
	configData.DbConfig.MaxLifeTime = dbMaxLifeTime
	configData.DbConfig.MaxLifeTime = dbMaxLifeTime
	configData.DbConfig.LogMode = logMode

	return configData, nil
}

func Run() {
	// Adding zerolog
	zerolog.TimeFieldFormat = "02-01-2006 15:04:05"
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// load config di dalam env
	configData, err := initEnv()
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Info().Msg(fmt.Sprintf("config data %v", configData))

	conn, err := config.Connect(configData, log.Logger)
	if err != nil {
		log.Error().Msg("RunService.NewPostgreSql.err : " + err.Error())
	}

	maxLifeTime, err := time.ParseDuration(configData.DbConfig.MaxLifeTime)
	if err != nil {
		log.Error().Msg("RunService.duration.err : " + err.Error())
		return
	}

	conn.SetConnMaxLifetime(maxLifeTime)
	conn.SetMaxIdleConns(configData.DbConfig.MaxIdle)
	conn.SetMaxOpenConns(configData.DbConfig.MaxConn)

	defer func() {
		errClose := conn.Close()
		if errClose != nil {
			log.Error().Msg(errClose.Error())
		}
	}()

	// setup timezone
	time.Local = time.FixedZone("Asia/Jakarta", 7*60*60)
	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  false,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           120 * time.Second,
	}))

	// open file app.log
	now := time.Now().Format("2006-01-02")
	logDir := "logger/"
	logFileName := now + " logger.log"
	logFilePath := logDir + logFileName

	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.MkdirAll(logDir, 0755); err != nil {
			log.Fatal().Err(err).Msg("Unable to create directory")
		}
	}

	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to open log file")
	}
	defer file.Close()

	// config logger zerolog
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: file}).With().Caller().Logger()

	r.Use(logger.SetLogger(
		logger.WithLogger(func(_ *gin.Context, l zerolog.Logger) zerolog.Logger {
			return l.Output(file).With().Logger()
		}),
	))

	// gin recovery for handle panic
	r.Use(gin.Recovery())

	initializeDomainModule(r, conn)

	version := "0.0.1"
	log.Info().Msg(fmt.Sprintf("Service Running version %s", version))
	addr := flag.String("port: ", ":"+configData.AppConfig.Port, "Address to listen and serve")
	err = r.Run(*addr)
	if err != nil {
		log.Error().Msg(err.Error())
	}
}

func initializeDomainModule(r *gin.Engine, db *sql.DB) {
	apiGroup := r.Group("/api")
	v1Group := apiGroup.Group("/v1")

	router.InitRoute(v1Group, db)
}
