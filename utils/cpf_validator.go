
package utils

import "regexp"

func CPFValid(cpf string) bool {
	cpf = regexp.MustCompile(`\D`).ReplaceAllLiteralString(cpf,"")
	if len(cpf) != 11 {
		return false
	}
	todosIguais := true
	for i := 1; i < 11; i++ {
		if cpf[i] != cpf[0] {
			todosIguais = false
			break
		}
	}
	if todosIguais {
		return false
	}
	d1, d2 := calcularDigitosVerificadores(cpf)

	return int(cpf[9]-'0') == d1 && int(cpf[10]-'0') == d2
}

func calcularDigitosVerificadores(cpf string) (int, int) {
	soma := 0
	for i := 0; i < 9; i++ {
		soma += int(cpf[i]-'0') * (10 - i)
	}
	resto := soma % 11
	d1 := 11 - resto
	if d1 >= 10 {
		d1 = 0
	}

	soma = 0
	for i := 0; i < 10; i++ {
		soma += int(cpf[i]-'0') * (11 - i)
	}
	resto = soma % 11
	d2 := 11 - resto
	if d2 >= 10 {
		d2 = 0
	}
	return d1, d2
}