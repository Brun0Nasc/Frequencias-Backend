package frequencias

import (
	"fmt"
	"net/http"

	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/frequencias/model"
	"github.com/gin-gonic/gin"
	"github.com/Brun0Nasc/Frequencias-Backend/domain/frequencias"
)

func NovaFrequencia(c *gin.Context) {
	fmt.Println("Tentando cadastrar uma nova frequencia")
	req := modelApresentacao.ReqFrequencia{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Could not create. Parameters were not passed correctly",
			"err":err.Error(),
		})
		return
	}

	err := frequencias.NovaFrequencia(&req)
	 
	if err != nil{
		c.JSON(400, gin.H{"err":err.Error()})
		return
	}
	c.JSON(201, gin.H{"message":"Frequencia adicionada com sucesso!"})
}