package utils

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Class-Connect-GRUPO-5/microservices-common/logger"
)

// PinGenerator interface defines the method for generating a PIN
type PinGenerator interface {
	GeneratePin() string
}

// RandomPinGenerator is the struct that implements the PinGenerator interface
type RandomPinGenerator struct{}

func NewRandomPinGenerator() RandomPinGenerator {
	return RandomPinGenerator{}
}

// GeneratePin generates a 6-digit random authentication PIN as a string.
// It uses a pseudo-random number generator seeded with the current Unix timestamp in nanoseconds.
// The generated PIN is zero-padded to ensure it is always 6 digits long.
// A debug log is created with the generated PIN for tracking purposes.
// Returns:
//
//	string: A 6-digit random PIN.
func (r *RandomPinGenerator) GeneratePin() string {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	pin := fmt.Sprintf("%06d", rng.Intn(1000000))

	logger.Logger.Debugf("Generated Authentication PIN: %s", pin)

	return pin
}
