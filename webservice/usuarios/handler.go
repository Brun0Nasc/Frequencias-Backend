package usuarios

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/usuarios/model"
	"github.com/Brun0Nasc/Frequencias-Backend/domain/usuarios"
)

func pegaJSON(c *gin.Context) *modelApresentacao.ReqUsuario {
	var req = modelApresentacao.ReqUsuario{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message":"Could not create. Parameters were not passed correctly",
			"err":err.Error(),
		})
		return nil
	}
	return &req
}

func novoUsuario(c *gin.Context){
	fmt.Println("Adicionando novo usuario")

	req := pegaJSON(c)

	err := usuarios.NovoUsuario(req)

	if err != nil{
		c.JSON(400, gin.H{"err":err.Error()})
		return
	}

	c.JSON(201, gin.H{"message":"Usuário adicionado com sucesso!"})
}

func listarUsuarios(c *gin.Context) {
	fmt.Println("Tentando listar usuarios")
	order := c.Param("order")

	intOrder, err := strconv.Atoi(order)
	if err != nil {
		c.JSON(404, gin.H{"message":"Parâmetro passado de forma incorreta", "err":err.Error()})
		return
	}

	res, err := usuarios.ListarUsuarios(intOrder)
	if err != nil {
		c.JSON(404, gin.H{"err":err.Error()})
		return
	}

	c.JSON(200, res)
}