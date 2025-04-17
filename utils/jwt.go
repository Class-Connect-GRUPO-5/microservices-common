package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func ParseJWT(tokenString string, jwtSecret string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil
}

// GenerateJWT generates a JSON Web Token (JWT) for a user with the specified user ID and role.
// The token is signed using the provided secret string and is set to expire in 1 hour.
//
// Parameters:
//   - userID: A string representing the unique identifier of the user.
//   - role: A string representing the role of the user.
//   - secretString: A string used as the secret key to sign the token.
//
// Returns:
//   - A string containing the signed JWT.
//   - An error if there is an issue during the token generation or signing process.
func GenerateJWT(userID string, role string, secretString string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour).Unix()
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     expirationTime,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := []byte(secretString)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	return signedToken, nil
}
