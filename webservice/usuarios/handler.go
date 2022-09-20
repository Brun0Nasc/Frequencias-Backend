package usuarios

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func novoUsuario(c *gin.Context){
	fmt.Println("Adicionando novo usuario")
}