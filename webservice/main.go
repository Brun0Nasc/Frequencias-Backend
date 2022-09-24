package main

import (
	"log"

	"os"

	"github.com/Brun0Nasc/Frequencias-Backend/config/server/middlewares"
	sync "github.com/Brun0Nasc/Frequencias-Backend/services/sync"
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

	// Inicializando rotinas
	if erro := sync.IniciarRotinas(); erro != nil {
		log.Fatal(erro)
	}

	r.Run(":" + port)
}
