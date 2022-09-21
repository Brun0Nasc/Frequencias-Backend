package main

import (
	"github.com/Brun0Nasc/Frequencias-Backend/config/server/middlewares"
	"github.com/Brun0Nasc/Frequencias-Backend/webservice/frequencias"
	"github.com/Brun0Nasc/Frequencias-Backend/webservice/login"
	"github.com/Brun0Nasc/Frequencias-Backend/webservice/usuarios"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	us := r.Group("usuarios", middlewares.Auth())
	fr := r.Group("frequencias", middlewares.Auth())
	lo := r.Group("login")

	frequencias.Router(fr)
	usuarios.Router(us)
	login.Router(lo)

	r.Run("localhost:3030")
}