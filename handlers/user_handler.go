package handlers

import(
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"api-golang/models"
	"api-golang/utils"
)

func CadastrarUsuario(c *gin.Context){
	var user models.User

	if err := c.ShouldBindJSON(&user); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(user.Nome) < 3 || strings.ContainsAny(user.Nome, "0123456789") || len (user.Sobrenome) < 3 || strings.ContainsAny(user.Sobrenome, "0123456789"){
		c.JSON(http.StatusBadRequest, gin.H{"error":"Nome e sobrenome devem ter mais de 3 caracteres e não possuir números"})
		return
	}

	if !utils.CPFValid(user.CPF) {
		c.JSON(http.StatusBadRequest, gin.H{"error":"CPF inválido"})
		return
	}

	if !utils.TelefoneValid(user.Telefone) {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Telefone inválido"})
		return
	}

	if !utils.EmailValid(user.Email){
		c.JSON(http.StatusBadRequest, gin.H{"error":"E-mail inválido"})
	}

	c.JSON(http.StatusCreated, gin.H{
		"menssage": "Usuário cadastrado!",
		"user": user,
	})
}