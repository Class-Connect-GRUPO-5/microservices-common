package middleware

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/Class-Connect-GRUPO-5/microservices-common/logger"
	"github.com/Class-Connect-GRUPO-5/microservices-common/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// RequireRole is a middleware that checks if the user has the required role.
func RequireRole(jwtSecret string, isIDRequired bool, requiredRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := ExtractUserJWT(c, jwtSecret)
		if err != nil {
			logger.Logger.Warnf("Error extracting JWT: %v", err)
			utils.HandleSuccess(c, http.StatusUnauthorized, err.Error(), nil)
			return
		}

		role := claims["role"]

		for _, requiredRole := range requiredRoles {
			if role == requiredRole {

				if isIDRequired {
					userID := claims["user_id"]
					reqUserID := c.Param("id_user")
					if userID != reqUserID {
						logger.Logger.Warnf("User ID mismatch: token user ID %v, request user ID %v", userID, reqUserID)
						utils.HandleSuccess(c, http.StatusForbidden, "User ID mismatch.", nil)
						return
					}
				}

				c.Set("user_id", claims["user_id"])
				c.Set("role", role)
				logger.Logger.Debugf("User %v has role %v", claims["user_id"], role)
				c.Next()
				return
			}
		}

		logger.Logger.Warnf("Access denied for user %v with role %v", claims["user_id"], role)
		utils.HandleSuccess(c, http.StatusForbidden, "Access denied.", nil)

		exp, ok := claims["exp"].(int64)
		if !ok {
			logger.Logger.Warnf("Invalid token expiration format for user %v", claims["user_id"])
			utils.HandleSuccess(c, http.StatusInternalServerError, "Invalid token expiration format.", nil)
			return
		} else if time.Now().Unix() > int64(exp) {
			logger.Logger.Warnf("Token expired for user %v", claims["user_id"])
			utils.HandleSuccess(c, http.StatusUnauthorized, "Token expired.", nil)
			return
		}

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

// GenerateAdminToken generates a JWT token for an admin user with the specified admin ID and role.
// The token includes claims for the user ID, role, and an expiration time of 24 hours.
// It uses the provided JWT secret to sign the token.
//
// Parameters:
//   - adminID: A string representing the unique identifier of the admin user.
//   - role: A string representing the role of the admin user.
//   - jwtSecret: A string used as the secret key to sign the JWT token.
//
// Returns:
//   - A string containing the signed JWT token.
//   - An error if the token signing process fails.
func GenerateAdminToken(adminID, role string, jwtSecret string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": adminID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 24hs
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
