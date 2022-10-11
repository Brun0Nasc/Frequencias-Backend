package main

import (
	"log"
	"os"

	"github.com/Brun0Nasc/Frequencias-Backend/config/server/middlewares"
	"github.com/Brun0Nasc/Frequencias-Backend/docs"
	"github.com/Brun0Nasc/Frequencias-Backend/services/sync"
	"github.com/Brun0Nasc/Frequencias-Backend/webservice/frequencia_usuario"
	"github.com/Brun0Nasc/Frequencias-Backend/webservice/login"
	"github.com/Brun0Nasc/Frequencias-Backend/webservice/usuarios"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @contact.name   Equipe Komanda
// @contact.url    http://www.swagger.io/support
// @contact.email  caiosousafernandesferreira@hotmail.com

// @license.name  Mozilla Public License 2.0
// @license.url   https://www.mozilla.org/en-US/MPL/2.0/
// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {
	port := os.Getenv("PORT")
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())

	docs.SwaggerInfo.Title = "Sistema de Frequencia"
	docs.SwaggerInfo.Description = "REST API Desafio"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "app-ponto-aprendizes.herokuapp.com"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"https"}

	usuarios.Router(r.Group("usuarios", middlewares.Auth()))
	frequencia_usuario.Router(r.Group("frequencia", middlewares.Auth()))
	login.Router(r.Group("login")) //! nada de middlewares aqui

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Inicializando rotinas
	if erro := sync.IniciarRotinas(); erro != nil {
		log.Fatal(erro)
	}

	r.Run(":" + port)
	//r.Run()
}
