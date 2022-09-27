package usuarios

import (
	"fmt"

	"github.com/Brun0Nasc/Frequencias-Backend/config/database"
	"github.com/Brun0Nasc/Frequencias-Backend/config/services"
	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/usuarios/model"
	"github.com/Brun0Nasc/Frequencias-Backend/infra/usuarios"
	"github.com/Brun0Nasc/Frequencias-Backend/utils"
)

func NovoUsuario(req *modelApresentacao.Usuario) (err error) {
	db := database.Conectar()
	defer db.Close()
	usuariosRepo := usuarios.NovoRepo(db)

	req.Senha = services.SHAR256Encoder(req.Senha)
	err = usuariosRepo.NovoUsuario(req)
	if err != nil {
		return fmt.Errorf("usuario nao adicionado // " + err.Error())
	}

	return
}

func ListarUsuarios(params *utils.RequestParams) (res *modelApresentacao.ListaUsuarios, err error) {
	db := database.Conectar()
	defer db.Close()
	usuariosRepo := usuarios.NovoRepo(db)

	res, err = usuariosRepo.ListarUsuarios(params)
	if err != nil {
		return nil, fmt.Errorf("usuarios nao listados // " + err.Error())
	}

	return
}

func InativarUsuario(id int) (err error) {
	db := database.Conectar()
	defer db.Close()
	usuariosRepo := usuarios.NovoRepo(db)

	err = usuariosRepo.InativarUsuario(id)
	if err != nil {
		return fmt.Errorf("usuario nao inativado // " + err.Error())
	}

	return
}
