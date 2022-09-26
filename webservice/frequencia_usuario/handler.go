package frequencia_usuario

import (
	"fmt"

	"github.com/Brun0Nasc/Frequencias-Backend/domain/frequencia_usuario"
	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/frequencia_usuario/model"
	//"github.com/Brun0Nasc/Frequencias-Backend/utils"
	"github.com/gin-gonic/gin"
)

func PegaJSON(c *gin.Context) *modelApresentacao.ReqFrequenciaUsuario {
	var req = modelApresentacao.ReqFrequenciaUsuario{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Could not create. Parameters were not passed correctly",
			"err":     err.Error(),
		})
		return nil
	}
	return &req
}

func NovaFrequenciaUsuario(c *gin.Context) {
	fmt.Println("Registrando Frequencia")

	req := PegaJSON(c)

	err := frequencia_usuario.NovaFrequenciaUsuario(req)

	if err != nil {
		c.JSON(400, gin.H{"err":err.Error()})
		return
	}

	c.JSON(201, gin.H{"message":"Frequência registrada"})
}