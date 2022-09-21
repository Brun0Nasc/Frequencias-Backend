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
		return fmt.Errorf("usuario nao adicionado // "+err.Error())
	}

	return
}

func ListarUsuarios(order int) (res []modelApresentacao.ReqUsuario, err error) {
	db := database.Conectar()
	defer db.Close()
	usuariosRepo := usuarios.NovoRepo(db)

	res, err = usuariosRepo.ListarUsuarios(order)
	if err != nil {
		return nil, fmt.Errorf("usuarios nao listados // "+err.Error())
	}

	return
}

func InativarUsuario(id int) (err error) {
	db := database.Conectar()
	defer db.Close()
	usuariosRepo := usuarios.NovoRepo(db)

	err = usuariosRepo.InativarUsuario(id)
	if err != nil {
		return fmt.Errorf("usuario nao inativado // "+err.Error())
	}

	return
}