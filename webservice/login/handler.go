package login

import (
	"github.com/Brun0Nasc/Frequencias-Backend/config/services"

	"net/http"

	"github.com/Brun0Nasc/Frequencias-Backend/domain/login"
	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/login/model"

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
	req := modelApresentacao.Login{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Could not create. Parameters were not passed correctly",
			"err":     err.Error(),
		})
		return
	}

	user, err := login.LoginUsuario(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if user == nil && err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Credentials",
		})
		return
	}

	// Checks if there is an error in this request
	token, err := services.NewJWTService().GenerateToken(user.ID, user.Tipo, user.Nome)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// If everything is true the token is generated
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
