package main

import (
	"log"

	sync "github.com/Brun0Nasc/Frequencias-Backend/services/sync"
	"github.com/Brun0Nasc/Frequencias-Backend/webservice/frequencias"
	"github.com/Brun0Nasc/Frequencias-Backend/webservice/usuarios"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	us := r.Group("usuarios")
	fr := r.Group("frequencias")

	frequencias.Router(fr)
	usuarios.Router(us)

	// Inicializando rotinas
	if erro := sync.IniciarRotinas(); erro != nil {
		log.Fatal(erro)
	}

	r.Run("localhost:3030")
}
