package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"time"
	_ "time"

	"api-golang/models"
	"api-golang/utils"
)

func CadastrarUsuario(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(user.Nome) < 3 || strings.ContainsAny(user.Nome, "0123456789") || len(user.Sobrenome) < 3 || strings.ContainsAny(user.Sobrenome, "0123456789") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nome e sobrenome devem ter mais de 3 caracteres e não possuir números"})
		return
	}

	if !utils.CPFValid(user.CPF) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CPF inválido"})
		return
	}

	if !utils.TelefoneValid(user.Telefone) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Telefone inválido"})
		return
	}

	if !utils.EmailValid(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "E-mail inválido"})
	}

	dataNascimento, err := time.Parse("02/01/2006", user.DataNascimento)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de data de nascimento inválido (DD/MM/AAAA)"})
	}

	imc := calcularIMC(user.Peso, user.Altura)

	classificacaoIMC := classificarIMC(imc)

	nomeCompleto := user.Nome + " " + user.Sobrenome

	idade := CalcularIdade(dataNascimento)

	faixaEtaria := calcularFaixaEtaria(idade)

	cpfFormatado := formatarCPF(user.CPF)
	telefoneFormatado := formatarTelefone(user.Telefone)

	id := uuid.New().String()

	c.JSON(http.StatusCreated, gin.H{
		"menssage":         "Usuário cadastrado!",
		"id":				id,
		"user":             user,
		"imc":              imc,
		"idade":            idade,
		"cpf":				cpfFormatado,
		"telefone":			telefoneFormatado,
		"faixaEtaria":      faixaEtaria,
		"classificacaoIMC": classificacaoIMC,
		"nomeCompleto":     nomeCompleto,
	})
}

func calcularIMC(peso float64, altura float64) float64 {
	return peso / (altura * altura)
}

func classificarIMC(imc float64) string {
	if imc < 18.5 {
		return "Magreza"
	}
	if imc >= 18.5 && imc <= 24.9 {
		return "Normal"
	}
	if imc >= 25 && imc <= 29.9 {
		return "Sobrepeso"
	}
	if imc > 30 {
		return "Obesidade"
	}
	return ""
}

func CalcularIdade(dataNascimento time.Time) int {
	hoje := time.Now()
	diferenca := hoje.Sub(dataNascimento)
	return int(diferenca.Hours() / 24 / 365)
}

func calcularFaixaEtaria(idade int) string {
	if idade < 12 {
		return "Criança"
	}
	if idade >= 12 && idade <= 17 {
		return "jovem"
	}
	if idade >= 18 && idade <= 59 {
		return "Adulto"
	}
	if idade >= 60 {
		return "idoso"
	}
	return ""
}

func formatarCPF(cpf string) string {
	cpfLimpo := strings.ReplaceAll(cpf, ".", "")
	cpfLimpo = strings.ReplaceAll(cpfLimpo, "-", "")
	if len(cpfLimpo) != 11 {
		return cpf
	}
	return cpfLimpo[:3] + "." + cpfLimpo[3:6] + "." + cpfLimpo[6:9] + "-" + cpfLimpo[9:]
}

func formatarTelefone(telefone string) string {
	telefoneLimpo := strings.ReplaceAll(telefone, "(", "")
	telefoneLimpo = strings.ReplaceAll(telefoneLimpo, ")", "")
	telefoneLimpo = strings.ReplaceAll(telefoneLimpo, " ", "")
	telefoneLimpo = strings.ReplaceAll(telefoneLimpo, "-", "")
	if len(telefoneLimpo) < 10 || len(telefoneLimpo) > 11 {
		return telefone
	}
	if len(telefoneLimpo) == 10 {
		return "(" + telefoneLimpo[:2] + ") " + telefoneLimpo[2:6] + "-" + telefoneLimpo[6:]
	}
	return "(" + telefoneLimpo[:2] + ") " + telefoneLimpo[2:7] + "-" + telefoneLimpo[7:]
}