package models

type User struct {
	Nome	string `json:"nome"`
	Sobrenome	string `json:"sobrenome"`
	DataNascimento string `json:"data_nascimento"`
	CPF string `json:"cpf"`
	Telefone string `json:"telefone"`
	Altura float64 `json:"altura"`
	Peso float64 `json:"peso"`
	Email string `json:"email"`
}