package frequencia_usuario

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.POST("/", NovaFrequenciaUsuario)
	r.GET("/lista", ListaFrequenciasData)
}