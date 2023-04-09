package pkg

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	minPasswordLength = 8  // Minimum length for a password.
	maxPasswordLength = 72 // Maximum length for a password.
)

// generateRandomString generates a random string of the given length.
// It returns an error if the random number generator fails.
func generateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}

// HashPassword generates a salted hash for the given password.
// It returns an error if the password is too short or too large, or if the salt or hash generation fails.

// ComparePasswordAndHash compares a password to a salted hash.
// It returns an error if the comparison fails.
func ComparePasswordAndHash(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return fmt.Errorf("invalid password: %w", err)
	}
	return nil
}
