package encrypt

import (
	"crypto/sha256"
	"encoding/hex"
)

func EncrytpHash(senha string) string {
	senhaB := []byte(senha)
	//senhaH := sha256.New().Sum(senhaB)
	senhaH := sha256.New()
	senhaH.Write(senhaB)
	senhaH.Sum(nil)
	senhaS := hex.EncodeToString(senhaB)
	return senhaS
}
