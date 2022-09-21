package login

import (
	"github.com/Brun0Nasc/Frequencias-Backend/config/services"
	
	"net/http"


	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/login/model"
	"github.com/Brun0Nasc/Frequencias-Backend/domain/login"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	req := modelApresentacao.Login{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message":"Could not create. Parameters were not passed correctly",
			"err":err.Error(),
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
			"error": "Invalid Credentials: ",
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
