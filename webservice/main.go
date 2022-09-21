package main

import (
	"log"

	"os"

	"github.com/Brun0Nasc/Frequencias-Backend/config/server/middlewares"
	sync "github.com/Brun0Nasc/Frequencias-Backend/services/sync"
	"github.com/Brun0Nasc/Frequencias-Backend/webservice/frequencias"
	"github.com/Brun0Nasc/Frequencias-Backend/webservice/login"
	"github.com/Brun0Nasc/Frequencias-Backend/webservice/usuarios"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	r := gin.Default()
	r.Use(cors.Default())

	usuarios.Router(r.Group("usuarios"))
	login.Router(r.Group("login", middlewares.Auth()))
	frequencias.Router(r.Group("frequencias", middlewares.Auth()))

	// Inicializando rotinas
	if erro := sync.IniciarRotinas(); erro != nil {
		log.Fatal(erro)
	}

	r.Run(":" + port)
}
