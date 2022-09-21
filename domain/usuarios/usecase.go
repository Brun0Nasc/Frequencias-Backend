package usuarios

import (
	"fmt"

	"github.com/Brun0Nasc/Frequencias-Backend/config/database"
	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/usuarios/model"
	"github.com/Brun0Nasc/Frequencias-Backend/infra/usuarios"
	"github.com/Brun0Nasc/Frequencias-Backend/config/services"
)

func NovoUsuario(req *modelApresentacao.ReqUsuario) (err error) {
	db := database.Conectar()
	defer db.Close()
	usuariosRepo := usuarios.NovoRepo(db)

	req.Senha = services.SHAR256Encoder(req.Senha)
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