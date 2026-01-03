package validation

import "unicode"

func ValidarCPF(cpf string) bool {
	// Remove caracteres não numéricos
	var digits []int
	for _, r := range cpf {
		if unicode.IsDigit(r) {
			digits = append(digits, int(r-'0'))
		}
	}

	// Deve ter 11 dígitos
	if len(digits) != 11 {
		return false
	}

	// Rejeita CPFs com todos os dígitos iguais
	allEqual := true
	for i := 1; i < 11; i++ {
		if digits[i] != digits[0] {
			allEqual = false
			break
		}
	}
	if allEqual {
		return false
	}

	// Primeiro dígito verificador
	sum := 0
	for i := 0; i < 9; i++ {
		sum += digits[i] * (10 - i)
	}
	d1 := (sum * 10) % 11
	if d1 == 10 {
		d1 = 0
	}
	if d1 != digits[9] {
		return false
	}

	// Segundo dígito verificador
	sum = 0
	for i := 0; i < 10; i++ {
		sum += digits[i] * (11 - i)
	}
	d2 := (sum * 10) % 11
	if d2 == 10 {
		d2 = 0
	}
	if d2 != digits[10] {
		return false
	}

	return true
}
