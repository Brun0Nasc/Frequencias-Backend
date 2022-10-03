package frequencia_usuario

import (
	"fmt"

	"github.com/Brun0Nasc/Frequencias-Backend/domain/frequencia_usuario"
	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/frequencia_usuario/model"
	"github.com/Brun0Nasc/Frequencias-Backend/utils"

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

// @Security bearerAuth
// @Summary POST uma nova frequencia
// @Description POST um novo usuario. O unico parametro que deve ser passado é o id do usuario em questão
// @Param		IDUsuario		body	string		true	"IDUsuario"
// @Accept json
// @Produce json
// @Success 201  {string} string "OK"
// @Failure 401,400  {string} string "error"
// @Tags Frequencias
// @Router /frequencia [post]
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

// @Security bearerAuth
// @Summary Listar Frequencias de um dia
// @Description Lista todas as frequencias de todas as pessoas, de um dia especifico.
// @Param		data		query			string				false		"data"		
// @Param		orderBy		query			string				false		"OrderBy" 		Enums(nome, entrada, saida)
// @Param		order		query			string				false		"order"			Enums(asc, desc)			
// @Accept json
// @Produce json
// @Success 201  {array} modelApresentacao.ListaUsuarioFrequencia "OK"
// @Failure 401,400  {string} string "error"
// @Tags Frequencias
// @Router /frequencia/lista [get]
func ListaFrequenciasData(c *gin.Context) {
	params := utils.ParseParams(c)

	res, err := frequencia_usuario.ListaFrequenciasData(params)
	if err != nil {
		c.JSON(404, gin.H{"err":err.Error()})
		return
	}

	c.JSON(200, res)
}