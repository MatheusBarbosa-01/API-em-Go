package handlers

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"api-golang/models"
)

func CadastrarUsuario(c *gin.Context){
	var user models.User

	if err := c.ShouldBindJSON(&user); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"menssage": "Usuário cadastrado!",
		"user": user,
	})
}