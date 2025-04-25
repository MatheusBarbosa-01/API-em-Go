
package utils

import "regexp"

func CPFValid(cpf string) bool {
	cpf = regexp.MustCompile(`\D`).ReplaceAllLiteralString(cpf,"")
	if len(cpf) != 11 {
		return false
	}
	//verificar uma biblioteca para fazer uma validação melhor
	return true
}