package usuarios

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.POST("/", novoUsuario)
	r.GET("/list_user", listarUsuarios)
	r.DELETE("/inativar/:id", inativarUsuario)
}