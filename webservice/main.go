package main

import (
	"os"

	"github.com/Brun0Nasc/Frequencias-Backend/config/server/middlewares"
	"github.com/Brun0Nasc/Frequencias-Backend/webservice/frequencias"
	"github.com/Brun0Nasc/Frequencias-Backend/webservice/login"
	"github.com/Brun0Nasc/Frequencias-Backend/webservice/usuarios"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())

	frequencias.Router(r.Group("frequencias", middlewares.Auth()))
	usuarios.Router(r.Group("usuarios", middlewares.Auth()))
	login.Router(r.Group("login")) //! nada de middlewares aqui

	r.Run(":" + port)
}
