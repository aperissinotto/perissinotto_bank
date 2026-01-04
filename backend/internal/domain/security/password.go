package security

import "golang.org/x/crypto/bcrypt"

func HashSenha(senhaAberta string) (string, bool) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(senhaAberta),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", false
	}
	return string(hash), true
}

func CompararSenha(senhaAberta, senhaHash string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(senhaHash),
		[]byte(senhaAberta),
	)

	if err != nil {
		return false
	}

	return true
}
