package test

import (
	"testing"
	"time"

	"github.com/Class-Connect-GRUPO-5/microservices-common/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

// Test for ParseJWT
func TestParseJWT_ValidToken(t *testing.T) {
	// Test secret
	secret := "test-secret-key"

	// Create a valid token for testing
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   "123",
		"role":      "admin",
		"email":     "test@example.com",
		"user_name": "Test User",
		"exp":       time.Now().Add(1 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	assert.NoError(t, err)

	// Parse the token
	claims, err := utils.ParseJWT(tokenString, secret)

	// Verify results
	assert.NoError(t, err)
	assert.NotNil(t, claims)
	assert.Equal(t, "123", claims["user_id"])
	assert.Equal(t, "admin", claims["role"])
	assert.Equal(t, "test@example.com", claims["email"])
	assert.Equal(t, "Test User", claims["user_name"])
}

func TestParseJWT_InvalidSignature(t *testing.T) {
	secret := "test-secret-key"

	// Create a token with correct structure but sign it with a different secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": "123",
		"exp":     time.Now().Add(1 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte("wrong-secret"))
	assert.NoError(t, err)

	// Parse with the correct secret - should fail
	claims, err := utils.ParseJWT(tokenString, secret)

	assert.Error(t, err)
	assert.Nil(t, claims)
}

func TestParseJWT_ExpiredToken(t *testing.T) {
	secret := "test-secret-key"

	// Create an expired token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": "123",
		"exp":     time.Now().Add(-1 * time.Hour).Unix(), // Expired 1 hour ago
	})

	tokenString, err := token.SignedString([]byte(secret))
	assert.NoError(t, err)

	// Parse the token
	claims, err := utils.ParseJWT(tokenString, secret)

	assert.Error(t, err)
	assert.Nil(t, claims)
}

func TestParseJWT_MalformedToken(t *testing.T) {
	secret := "test-secret-key"

	// Use a completely invalid token string
	tokenString := "not-a-valid-jwt-token"

	claims, err := utils.ParseJWT(tokenString, secret)

	assert.Error(t, err)
	assert.Nil(t, claims)
}

func TestParseJWT_UnexpectedSigningMethod(t *testing.T) {
	secret := "test-secret-key"

	// Create token with different signing method
	token := jwt.New(jwt.SigningMethodNone)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = "123"
	claims["exp"] = time.Now().Add(1 * time.Hour).Unix()

	tokenString, err := token.SignedString(jwt.UnsafeAllowNoneSignatureType)
	assert.NoError(t, err)

	parsedClaims, err := utils.ParseJWT(tokenString, secret)

	assert.Error(t, err)
	assert.Nil(t, parsedClaims)
}

// Tests for GenerateJWT
func TestGenerateJWT_Success(t *testing.T) {
	// Test parameters
	userID := "user123"
	role := "user"
	email := "user@example.com"
	userName := "Test User"
	secret := "test-secret-key"

	// Generate a token
	tokenString, err := utils.GenerateJWT(userID, role, email, userName, secret)

	// Verify token was generated
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)
}

func TestGenerateJWT_CanBeParsed(t *testing.T) {
	// Test parameters
	userID := "user123"
	role := "user"
	email := "user@example.com"
	userName := "Test User"
	secret := "test-secret-key"

	// Generate a token
	tokenString, err := utils.GenerateJWT(userID, role, email, userName, secret)
	assert.NoError(t, err)

	// Verify token can be parsed
	claims, err := utils.ParseJWT(tokenString, secret)
	assert.NoError(t, err)
	assert.NotNil(t, claims)
}

func TestGenerateJWT_ClaimsMatch(t *testing.T) {
	// Test parameters
	userID := "user123"
	role := "user"
	email := "user@example.com"
	userName := "Test User"
	secret := "test-secret-key"

	// Generate a token
	tokenString, err := utils.GenerateJWT(userID, role, email, userName, secret)
	assert.NoError(t, err)

	// Get claims and verify they match
	claims, err := utils.ParseJWT(tokenString, secret)
	assert.NoError(t, err)

	assert.Equal(t, userID, claims["user_id"])
	assert.Equal(t, role, claims["role"])
	assert.Equal(t, email, claims["email"])
	assert.Equal(t, userName, claims["user_name"])
}

func TestGenerateJWT_HasExpiration(t *testing.T) {
	// Test parameters
	userID := "user123"
	role := "user"
	email := "user@example.com"
	userName := "Test User"
	secret := "test-secret-key"

	// Generate a token
	tokenString, err := utils.GenerateJWT(userID, role, email, userName, secret)
	assert.NoError(t, err)

	// Get claims and verify expiration
	claims, err := utils.ParseJWT(tokenString, secret)
	assert.NoError(t, err)

	// Verify expiration time is in the future
	exp, ok := claims["exp"].(float64)
	assert.True(t, ok)
	assert.Greater(t, exp, float64(time.Now().Unix()))
}
