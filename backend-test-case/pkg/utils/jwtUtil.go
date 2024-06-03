package utils

import (
	"backend_test_case/model/dto"
	"backend_test_case/pkg/constants"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func GenerateJWT(id, username, role string) (string, error) {
	var (
		issuer = os.Getenv("JWT_ISSUER")
		method = jwt.SigningMethodHS256
	)

	if secretKey == nil {
		return "", errors.New(constants.ErrSecretKeyNotSet)
	}

	claims := dto.JWTClaims{
		ID:       id,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			Issuer:    issuer,
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(method, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string) (*jwt.Token, error) {
	claims := &dto.JWTClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func GetJWT(c *gin.Context) (*dto.JWTClaims, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, errors.New(constants.ErrAuthIsMissing)
	}

	tokenString := strings.ReplaceAll(authHeader, "Bearer ", "")

	token, err := VerifyJWT(tokenString)
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*dto.JWTClaims)
	return claims, nil
}
