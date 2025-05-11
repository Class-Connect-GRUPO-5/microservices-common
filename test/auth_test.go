package test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Class-Connect-GRUPO-5/microservices-common/logger"
	"github.com/Class-Connect-GRUPO-5/microservices-common/middleware"
	"github.com/Class-Connect-GRUPO-5/microservices-common/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	logger.InitLogger("auth", logger.Error, os.Stdout, true)
}

func setupGin() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	return r
}

func createAuthHeader(userID, role, email, userName, secret string) string {
	token, _ := utils.GenerateJWT(userID, role, email, userName, secret)
	return "Bearer " + token
}

func TestExtractUserJWT_ValidToken(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	secret := "test-secret"
	userID := "user123"
	role := "admin"
	email := "test@example.com"
	userName := "Test User"

	c.Request = httptest.NewRequest("GET", "/", nil)
	c.Request.Header.Set("Authorization", createAuthHeader(userID, role, email, userName, secret))

	claims, err := middleware.ExtractUserJWT(c, secret)

	assert.NoError(t, err)
	assert.NotNil(t, claims)
	assert.Equal(t, userID, claims["user_id"])
	assert.Equal(t, role, claims["role"])
	assert.Equal(t, email, claims["email"])
	assert.Equal(t, userName, claims["user_name"])
}

func TestExtractUserJWT_MissingHeader(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("GET", "/", nil)

	claims, err := middleware.ExtractUserJWT(c, "any-secret")

	assert.Error(t, err)
	assert.Nil(t, claims)
	assert.Contains(t, err.Error(), "authorization header missing")
}

func TestExtractUserJWT_InvalidToken(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("GET", "/", nil)
	c.Request.Header.Set("Authorization", "Bearer invalid.token.format")

	claims, err := middleware.ExtractUserJWT(c, "secret")

	assert.Error(t, err)
	assert.Nil(t, claims)
}

func TestRequireRole_Authorized(t *testing.T) {
	r := setupGin()
	secret := "test-secret"
	userID := "user123"
	role := "admin"

	var handlerCalled bool
	handler := func(c *gin.Context) {
		handlerCalled = true
		assert.Equal(t, userID, c.MustGet("user_id"))
		assert.Equal(t, role, c.MustGet("role"))
	}

	r.GET("/protected", middleware.RequireRole(secret, false, []string{"admin"}), handler)

	req := httptest.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", createAuthHeader(userID, role, "admin@test.com", "Admin", secret))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.True(t, handlerCalled)
}

func TestRequireRole_Unauthorized_WrongRole(t *testing.T) {
	r := setupGin()
	secret := "test-secret"

	var handlerCalled bool
	handler := func(c *gin.Context) {
		handlerCalled = true
	}

	r.GET("/protected", middleware.RequireRole(secret, false, []string{"admin"}), handler)

	req := httptest.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", createAuthHeader("user123", "user", "user@test.com", "User", secret))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.False(t, handlerCalled)
}

func TestRequireRole_Unauthorized_InvalidToken(t *testing.T) {
	r := setupGin()
	secret := "test-secret"

	var handlerCalled bool
	handler := func(c *gin.Context) {
		handlerCalled = true
	}

	r.GET("/protected", middleware.RequireRole(secret, false, []string{"admin"}), handler)

	req := httptest.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer invalid.token")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.False(t, handlerCalled)
}

func TestRequireRole_IDRequired_Match(t *testing.T) {
	r := setupGin()
	secret := "test-secret"
	userID := "user123"

	var handlerCalled bool
	handler := func(c *gin.Context) {
		handlerCalled = true
	}

	r.GET("/users/:id_user/profile", middleware.RequireRole(secret, true, []string{"user", "admin"}), handler)

	req := httptest.NewRequest("GET", "/users/user123/profile", nil)
	req.Header.Set("Authorization", createAuthHeader(userID, "user", "user@test.com", "User", secret))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.True(t, handlerCalled)
}

func TestRequireRole_IDRequired_Mismatch(t *testing.T) {
	r := setupGin()
	secret := "test-secret"

	var handlerCalled bool
	handler := func(c *gin.Context) {
		handlerCalled = true
	}

	r.GET("/users/:id_user/profile", middleware.RequireRole(secret, true, []string{"user", "admin"}), handler)

	req := httptest.NewRequest("GET", "/users/different-user/profile", nil)
	req.Header.Set("Authorization", createAuthHeader("user123", "user", "user@test.com", "User", secret))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.False(t, handlerCalled)
}

func TestRequireRole_MultipleAllowedRoles(t *testing.T) {
	r := setupGin()
	secret := "test-secret"

	var handlerCalled bool
	handler := func(c *gin.Context) {
		handlerCalled = true
	}

	r.GET("/protected", middleware.RequireRole(secret, false, []string{"user", "admin", "moderator"}), handler)

	req := httptest.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", createAuthHeader("user123", "moderator", "mod@test.com", "Moderator", secret))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.True(t, handlerCalled)
}
