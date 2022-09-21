package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", fmt.Errorf("failed to hash password: %w", err)
    }
    return string(hashedPassword), nil
}

func CheckPassword(password string, hashedPassword string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func TestPassword(t *testing.T) {
    password := RandomString(6)

    hashedPassword, err := HashPassword(password)
    require.NoError(t, err)
    require.NotEmpty(t, hashedPassword)

    err = CheckPassword(password, hashedPassword)
    require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}