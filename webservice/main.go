package main

import (
	"os"

	"github.com/Brun0Nasc/Frequencias-Backend/config/server/middlewares"
	"github.com/Brun0Nasc/Frequencias-Backend/webservice/frequencia_usuario"
	"github.com/Brun0Nasc/Frequencias-Backend/webservice/login"
	"github.com/Brun0Nasc/Frequencias-Backend/webservice/usuarios"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())

	frequencia_usuario.Router(r.Group("frequencia"))
	usuarios.Router(r.Group("usuarios"))
	login.Router(r.Group("login")) //! nada de middlewares aqui

	r.Run(":" + port)
}
