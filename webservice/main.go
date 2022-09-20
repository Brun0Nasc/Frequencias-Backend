package main

import (
	"github.com/Brun0Nasc/Frequencias-Backend/webservice/usuarios"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	us := r.Group("usuarios")

	usuarios.Router(us)

	r.Run("localhost:3030")
}