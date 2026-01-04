package security

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashSenha(senhaAberta string) (string, bool) {
	log.Printf("senha aberta (len=%d): '%s'\n", len(senhaAberta), senhaAberta)
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(senhaAberta),
		bcrypt.DefaultCost,
	)
	if err != nil {
		log.Println(err)
		return "", false
	}
	return string(hash), true
}

func CompararSenha(senhaAberta string, senhaHash string) bool {
	log.Printf("senha aberta (len=%d): '%s'\n", len(senhaAberta), senhaAberta)
	err := bcrypt.CompareHashAndPassword(
		[]byte(senhaHash),
		[]byte(senhaAberta),
	)

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
