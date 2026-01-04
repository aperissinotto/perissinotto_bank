package validation

import "regexp"

func ValidarSenha(senha string) bool {
	regex := `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`
	ok, err := regexp.MatchString(regex, senha)
	if err != nil {
		return false
	}
	return ok
}
