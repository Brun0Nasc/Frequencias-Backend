package frequencias

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.POST("/", NovaFrequencia)
	r.GET("/:id", ListarFrequenciasUsuario)
}
