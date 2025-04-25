package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"api-golang/handlers"
)

func main(){
	router := gin.Default()

	router.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})
	router.POST("/cadastro-usuario", handlers.CadastrarUsuario)
	router.Run(":8080")
}