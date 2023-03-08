package encrypt

import (
	"login-app/configuration"

	"golang.org/x/crypto/bcrypt"
)

func EncrytpHash(senha string, logs configuration.GoAppTools) string {
	senhaB := []byte(senha)
	senhaH, err := bcrypt.GenerateFromPassword(senhaB, 14)
	configuration.Check(err, logs)
	senhaS := string(senhaH)
	return senhaS
}
