package encrypt

import (
	"login-app/configuration"

	"golang.org/x/crypto/bcrypt"
)

func EncrytpHash(senha string, logs configuration.GoAppTools) string {
	senhaB := []byte(senha)
	senhaH, err := bcrypt.GenerateFromPassword(senhaB, bcrypt.MinCost)
	configuration.Check(err, logs)
	senhaS := string(senhaH)
	return senhaS
}

func CheckHash(senha string, hash string, logs configuration.GoAppTools) bool {
	senhaB := []byte(senha)
	hashB := []byte(hash)
	err := bcrypt.CompareHashAndPassword(hashB, senhaB)
	if err == nil {
		return true
	}
	configuration.Check(err, logs)
	return false
}
