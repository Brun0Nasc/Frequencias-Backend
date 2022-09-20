package usuarios

import (
	"fmt"

	"github.com/Brun0Nasc/Frequencias-Backend/config/database"
	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/usuarios/model"
	"github.com/Brun0Nasc/Frequencias-Backend/infra/usuarios"
)

func NovoUsuario(req *modelApresentacao.ReqUsuario) (err error) {
	db := database.Conectar()
	defer db.Close()
	usuariosRepo := usuarios.NovoRepo(db)

	err = usuariosRepo.NovoUsuario(req)
	if err != nil{
		return fmt.Errorf("user not added // "+err.Error())
	}

	return
}