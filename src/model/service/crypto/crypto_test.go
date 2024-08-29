package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashAndCheckPassword(t *testing.T) {
	crypto := &Crypto{}
	password := "securePassword"

	// HashPassword
	hashedPassword, err := crypto.HashPassword(password)
	assert.Nil(t, err, "An error was not expected when hashing the password, but it was returned")
	assert.NotEmpty(t, hashedPassword, "expected a non-empty hashed password but was returned empty")
	assert.NotEqual(t, password, hashedPassword, "the hashed password must not be the same as the original password")

	// CheckPassword com senha correta
	isValid, err := crypto.CheckPassword(password, hashedPassword)
	assert.Nil(t, err, "An error was not expected when checking the correct password, but it was returned")
	assert.True(t, isValid, "the hashed password should be valid when compared to the original password")

	// CheckPassword com senha incorreta
	isValid, err = crypto.CheckPassword("incorrectPassword", hashedPassword)
	assert.False(t, isValid, "Incorrect password must not be valid")
	assert.NotNil(t, err, "There must be an error for incorrect password")
	assert.Contains(t, err.Error(), "hashedPassword is not the hash of the given password", "The error should indicate that the password is incorrect")
}
