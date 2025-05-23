package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Class-Connect-GRUPO-5/microservices-common/logger"
	"github.com/Class-Connect-GRUPO-5/microservices-common/utils"
	"github.com/gin-gonic/gin"
)

// RequireRole is a middleware that checks if the user has the required role.
func RequireRole(jwtSecret string, isIDRequired bool, requiredRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := ExtractUserJWT(c, jwtSecret)
		if err != nil {
			logger.Logger.Warnf("Error extracting JWT: %v", err)
			utils.HandleError(c, http.StatusUnauthorized, "Unauthorized", err.Error())
			c.Abort()
			return
		}

		exp, ok := claims["exp"].(float64)
		if !ok {
			logger.Logger.Warnf("Invalid token expiration format for user %v", claims["user_id"])
			utils.HandleError(c, http.StatusBadRequest, "Invalid token expiration format.", fmt.Errorf("invalid token expiration format").Error())
			c.Abort()
			return
		} else if time.Now().Unix() > int64(exp) {
			logger.Logger.Warnf("Token expired for user %v", claims["user_id"])
			utils.HandleError(c, http.StatusUnauthorized, "Unauthorized", fmt.Errorf("token expired").Error())
			c.Abort()
			return
		}

		role := claims["role"]

		roleMatched := false
		for _, requiredRole := range requiredRoles {
			if role == requiredRole {
				roleMatched = true
				break
			}
		}

		if !roleMatched {
			logger.Logger.Warnf("Access denied for user %v with role %v", claims["user_id"], role)
			utils.HandleError(c, http.StatusUnauthorized, "Unauthorized", fmt.Errorf("access denied").Error())
			c.Abort()
			return
		}

		if isIDRequired {
			userID := claims["user_id"]
			reqUserID := c.Param("id_user")
			if userID != reqUserID {
				logger.Logger.Warnf("User ID mismatch: token user ID %v, request user ID %v", userID, reqUserID)
				utils.HandleError(c, http.StatusUnauthorized, "Unauthorized", fmt.Errorf("user ID mismatch").Error())
				c.Abort()
				return
			}
		}

		c.Set("user_id", claims["user_id"])
		c.Set("role", role)
		logger.Logger.Debugf("User %v with role %v authorized", claims["user_id"], role)
		c.Next()
	}
}

// ExtractUserJWT extracts the JWT from the request context and verifies it.
func ExtractUserJWT(c *gin.Context, jwtSecret string) (map[string]interface{}, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, errors.New("authorization header missing")
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := utils.ParseJWT(tokenString, jwtSecret)
	if err != nil {
		return nil, err
	}
	return claims, nil
}
