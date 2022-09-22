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

	us := r.Group("usuarios", middlewares.Auth())
	fr := r.Group("frequencias", middlewares.Auth())
	lo := r.Group("login") //! nada de middlewares aqui

	frequencias.Router(fr)
	usuarios.Router(us)
	login.Router(lo)

	r.Run(":" + port)
}
