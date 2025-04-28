package main

import (
	"api-golang/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})
	router.POST("/cadastro-usuario", handlers.CadastrarUsuario)
	router.Run(":8080")
}
