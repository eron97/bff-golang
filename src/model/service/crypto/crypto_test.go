package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashAndCheckPassword(t *testing.T) {
	crypto := &Crypto{}
	password := "senhaSegura123"

	// HashPassword
	hashedPassword, err := crypto.HashPassword(password)
	assert.Nil(t, err, "não era esperado um erro ao fazer hash da senha, mas foi retornado")
	assert.NotEmpty(t, hashedPassword, "esperava uma senha hash não vazia, mas foi retornada vazia")
	assert.NotEqual(t, password, hashedPassword, "a senha hash não deve ser igual à senha original")

	// CheckPassword com senha correta
	isValid, err := crypto.CheckPassword(password, hashedPassword)
	assert.Nil(t, err, "não era esperado um erro ao verificar a senha correta, mas foi retornado")
	assert.True(t, isValid, "a senha hash deveria ser válida quando comparada com a senha original")

	// CheckPassword com senha incorreta
	isValid, err = crypto.CheckPassword("senhaIncorreta", hashedPassword)
	assert.False(t, isValid, "A senha incorreta não deve ser válida")
	assert.NotNil(t, err, "Deve haver um erro para senha incorreta")
	assert.Contains(t, err.Error(), "hashedPassword is not the hash of the given password", "O erro deve indicar que a senha está incorreta")
}
