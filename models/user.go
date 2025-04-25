package models

type User struct {
	Nome	string `json:"nome" validate:"required,min=3,max=50"`
	Sobrenome	string `json:"sobrenome" validate:"required,min=3,max=50"`
	DataNascimento string `json:"data_nascimento" validate:"required"`
	CPF string `json:"cpf" validate:"required"` 
	Telefone string `json:"telefone" validate:"required"`
	Altura float64 `json:"altura" validate:"required,gt=0,lte=3"`
	Peso float64 `json:"peso" validate:"required,gt=0,lte=600"`
	Email string `json:"email" validate:"required,email"`
}