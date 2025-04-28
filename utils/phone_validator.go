package utils

import "regexp"

func TelefoneValid(telefone string) bool {
telefone = regexp.MustCompile(`\D`).ReplaceAllLiteralString(telefone,"")
return len(telefone) >=10 && len(telefone) <=11
}