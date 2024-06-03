package middleware

import (
	"backend_test_case/model/dto"
	"backend_test_case/model/dto/json"
	"backend_test_case/pkg/constants"
	"backend_test_case/pkg/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuth(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			json.NewResponseUnauthorized(c, "Invalid token", constants.AuthService, "01")
			c.Abort()
			return
		}

		tokenString := strings.ReplaceAll(authHeader, "Bearer ", "")
		token, err := utils.VerifyJWT(tokenString)
		if err != nil {
			json.NewResponseError(c, err.Error(), constants.AuthService, "02")
			c.Abort()
			return
		}

		if !token.Valid {
			json.NewResponseForbidden(c, "Forbidden", constants.AuthService, "03")
			c.Abort()
			return
		}
		claims := token.Claims.(*dto.JWTClaims)

		validRole := false
		if len(roles) > 0 {
			for _, role := range roles {
				if claims.Role == role {
					validRole = true
					break
				}
			}
		}

		if !validRole {
			json.NewResponseForbidden(c, "Forbidden", constants.AuthService, "04")
			c.Abort()
			return
		}
		c.Next()
	}
}
