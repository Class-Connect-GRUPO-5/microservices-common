package middleware

import (
	"fmt"
	"net/http"

	"github.com/Class-Connect-GRUPO-5/microservices-common/logger"
	"github.com/Class-Connect-GRUPO-5/microservices-common/utils"
	"github.com/gin-gonic/gin"
)

type UserData struct {
	UserID    string
	Role      string
	UserEmail string
	UserName  string
}

// SetJWTDataFromToken returns a middleware that extracts data from JWT token
// and sets it in the Gin context for subsequent handlers to use.
func SetJWTDataFromToken(secret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			logger.Logger.Warn("Authorization header is missing")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			return
		}

		var token string
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			token = authHeader[7:]
		} else {
			logger.Logger.Warn("Invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			return
		}

		logger.Logger.Debug("Parsing JWT token")
		claims, err := utils.ParseJWT(token, secret)
		if err != nil {
			logger.Logger.Errorf("Failed to parse token: %v", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			return
		}

		userID, ok := claims["user_id"].(string)
		if !ok {
			logger.Logger.Error("Missing or invalid user_id in token")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token structure"})
			return
		}

		role, ok := claims["role"].(string)
		if !ok {
			logger.Logger.Error("Missing or invalid role in token")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token structure"})
			return
		}

		email, ok := claims["email"].(string)
		if !ok {
			logger.Logger.Error("Missing or invalid email in token")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token structure"})
			return
		}

		name, ok := claims["user_name"].(string)
		if !ok {
			logger.Logger.Error("Missing or invalid name in token")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token structure"})
			return
		}

		jwtData := UserData{
			UserID:    userID,
			Role:      role,
			UserEmail: email,
			UserName:  name,
		}

		logger.Logger.Debug(fmt.Sprintf("Parsed JWT data: %+v", jwtData))

		ctx.Set("userData", jwtData)
		ctx.Next()
	}
}

// GetJWTUserData retrieves and validates user data from the JWT stored in the context.
// Returns the user data or an APIResponse error if not found/valid.
func GetJWTUserData(ctx *gin.Context) (UserData, error) {
	jwtData, exists := ctx.Get("userData")
	if !exists {
		return UserData{}, fmt.Errorf("Authorization error no tengo nati")
	}

	userData, ok := jwtData.(UserData)
	if !ok {
		return UserData{}, fmt.Errorf("Authorization error")
	}

	return userData, nil
}

// extractUserIDFromToken extracts the user ID from the JWT token in the Authorization header.
// Returns the user ID if successful, or an error if the token is invalid or missing.
func ExtractUserIDFromToken(ctx *gin.Context, secret string) (string, error) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		logger.Logger.Warn("Authorization header is missing")
		return "", fmt.Errorf("authorization header is required")
	}

	var token string
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		token = authHeader[7:]
	} else {
		logger.Logger.Warn("Invalid authorization header format")
		return "", fmt.Errorf("invalid authorization header format")
	}

	logger.Logger.Debug("Parsing JWT token")
	claims, err := utils.ParseJWT(token, secret)
	if err != nil {
		logger.Logger.Errorf("Failed to parse token: %v", err)
		return "", fmt.Errorf("failed to parse token: %w", err)
	}

	if claims["user_id"] == nil {
		logger.Logger.Errorf("Token missing user_id field. Claims: %+v", claims)
		return "", fmt.Errorf("invalid token: missing user_id field")
	}

	userID := claims["user_id"].(string)
	logger.Logger.Debug(fmt.Sprintf("User ID from token: %s", userID))

	return userID, nil
}
