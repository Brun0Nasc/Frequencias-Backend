package login

import (
	"github.com/Brun0Nasc/Frequencias-Backend/domain/login"
	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/usuarios/model"

	"github.com/gin-gonic/gin"
)

// Authenticate godoc
// @Summary Login para usuário
// @Description Autentica um usuário com as credenciais enviadas e se OK gera um token
// @ID Authentication
// @Accept json
// @Produce json
// @Param		Login		body	string		true	"Login"
// @Success 200 {string} string "OK"
// @Failure 400,500 {string} string "error"
// @Tags Usuarios
// @Router /login [post]
func Login(c *gin.Context) {
	req := modelApresentacao.Usuario{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Could not create. Parameters were not passed correctly",
			"err":     err.Error(),
		})
		return
	}

	token, err := login.LoginUsuario(&req)

	if err != nil {
		c.JSON(403, gin.H{"err": err.Error()})
		return
	}

	if token == "" && err == nil {
		c.JSON(403, gin.H{
			"error": err.Error(),
		})
		return
	}

	// If everything is true the token is generated
	c.JSON(200, gin.H{
		"token": token,
	})

}
