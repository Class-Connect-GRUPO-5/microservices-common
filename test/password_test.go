// filepath: /home/juan/Documents/facultad/is2/microservices-common/test/password_test.go
package test

import (
	"os"
	"testing"

	"github.com/Class-Connect-GRUPO-5/microservices-common/logger"
	"github.com/Class-Connect-GRUPO-5/microservices-common/utils"
	"github.com/stretchr/testify/assert"
)

func init() {
	logger.InitLogger("password_test", logger.Error, os.Stdout, true)
}

func TestHashPassword_Success(t *testing.T) {
	password := "MySecurePassword123!"

	hashedPassword, err := utils.HashPassword(password)

	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)
	assert.NotEqual(t, password, hashedPassword)
}

func TestHashPassword_EmptyPassword(t *testing.T) {
	password := ""

	hashedPassword, err := utils.HashPassword(password)

	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)
}

func TestCheckPassword_CorrectPassword(t *testing.T) {
	password := "MySecurePassword123!"

	hashedPassword, err := utils.HashPassword(password)
	assert.NoError(t, err)

	err = utils.CheckPassword(hashedPassword, password)

	assert.NoError(t, err)
}

func TestCheckPassword_IncorrectPassword(t *testing.T) {
	password := "MySecurePassword123!"
	hashedPassword, err := utils.HashPassword(password)
	assert.NoError(t, err)

	wrongPassword := "WrongPassword123!"
	err = utils.CheckPassword(hashedPassword, wrongPassword)

	assert.Error(t, err)
}

func TestCheckPassword_EmptyPassword(t *testing.T) {
	password := "MySecurePassword123!"
	hashedPassword, err := utils.HashPassword(password)
	assert.NoError(t, err)

	err = utils.CheckPassword(hashedPassword, "")

	assert.Error(t, err)
}

func TestCheckPassword_EmptyHash(t *testing.T) {
	err := utils.CheckPassword("", "somepassword")

	assert.Error(t, err)
}

func TestHashPassword_DifferentInputs_DifferentHashes(t *testing.T) {
	password1 := "MySecurePassword123!"
	password2 := "MySecurePassword123#"

	hash1, err1 := utils.HashPassword(password1)
	hash2, err2 := utils.HashPassword(password2)

	assert.NoError(t, err1)
	assert.NoError(t, err2)

	assert.NotEqual(t, hash1, hash2)
}

func TestHashPassword_SameInput_DifferentHashes(t *testing.T) {
	password := "MySecurePassword123!"

	hash1, err1 := utils.HashPassword(password)
	hash2, err2 := utils.HashPassword(password)

	assert.NoError(t, err1)
	assert.NoError(t, err2)

	assert.NotEqual(t, hash1, hash2)
}

func TestCrossCheck_HashAndVerify(t *testing.T) {
	passwords := []string{
		"Simple123",
		"Complex!@#$%^&*()",
		"WithSpaces And Symbols!",
		"1234567890",
		"abcdefghijklmnopqrstuvwxyz",
	}

	for _, password := range passwords {
		// Hash password
		hashedPassword, err := utils.HashPassword(password)
		assert.NoError(t, err)

		// Verify correct password works
		err = utils.CheckPassword(hashedPassword, password)
		assert.NoError(t, err, "Failed to verify correct password: %s", password)

		// Verify wrong password fails
		err = utils.CheckPassword(hashedPassword, password+"wrong")
		assert.Error(t, err, "Wrong password should fail: %s", password+"wrong")
	}
}
