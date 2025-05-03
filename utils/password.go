package utils

import (
	"github.com/Class-Connect-GRUPO-5/microservices-common/logger"

	"golang.org/x/crypto/bcrypt"
)

func CheckPassword(hashedPwd string, plainPwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
}

// HashPassword takes a plain text password as input and returns its hashed representation
// using the bcrypt algorithm. The function ensures that the password is securely hashed
// with a default cost factor.
//
// Parameters:
//   - password: The plain text password to be hashed.
//
// Returns:
//   - A string containing the hashed password.
//   - An error if the hashing process fails.
//
// Note:
//   - The bcrypt algorithm is computationally expensive, which helps protect against
//     brute-force attacks.
//   - Ensure that the returned error is handled appropriately to avoid security risks.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Logger.Errorf("error while hashing the password: %s", err)
		return "", err
	}

	return string(hashedPassword), nil
}
