package usuarios

import (
	"fmt"
	"strconv"

	"github.com/Brun0Nasc/Frequencias-Backend/domain/usuarios"
	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/usuarios/model"
	"github.com/Brun0Nasc/Frequencias-Backend/utils"
	"github.com/gin-gonic/gin"
)

func pegaJSON(c *gin.Context) *modelApresentacao.Usuario {
	var req = modelApresentacao.Usuario{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Could not create. Parameters were not passed correctly",
			"err":     err.Error(),
		})
		return nil
	}
	return &req
}

// @Security bearerAuth
// @Summary POST a new User
// @Description POST a new User. For the request to be met, the "nome", "email", "password", are required.
// @Param		NewUser		body	string		true	"NewUser"
// @Accept json
// @Produce json
// @Success 201  {string} string "OK"
// @Failure 401,400  {string} string "error"
// @Tags Usuarios
// @Router /usuarios [post]
func novoUsuario(c *gin.Context) {
	fmt.Println("Adicionando novo usuario")

	req := pegaJSON(c)

	err := usuarios.NovoUsuario(req)

	if err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Usuário adicionado com sucesso!"})
}

func listarUsuarios(c *gin.Context) {
	params := utils.ParseParams(c)

	res, err := usuarios.ListarUsuarios(params)
	if err != nil {
		c.JSON(404, gin.H{"err": err.Error()})
		return
	}

	c.JSON(200, res)
}

func inativarUsuario(c *gin.Context) {
	fmt.Println("Tentando inativar usuario")
	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(404, gin.H{"err": err.Error()})
		return
	}

	err = usuarios.InativarUsuario(intId)
	if err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Cadastro inativado com sucesso"})
}
