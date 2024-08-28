package crypto

import "golang.org/x/crypto/bcrypt"

type CryptoInterface interface {
	HashPassword(password string) (string, error)
	CheckPassword(plainPassword, hashedPassword string) (bool, error)
}

type Crypto struct{}

func (c *Crypto) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (c *Crypto) CheckPassword(plainPassword, hashedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		return false, err
	}
	return true, nil
}
