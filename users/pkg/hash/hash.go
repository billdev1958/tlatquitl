package hash

import "golang.org/x/crypto/bcrypt"

// HashPassword generates a bcrypt hash from a password string.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares a bcrypt password hash with its plaintext equivalent.
// Returns true if they match, false otherwise.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
