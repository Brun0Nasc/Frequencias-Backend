package usuarios

import (
	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/usuarios/model"
	"github.com/Brun0Nasc/Frequencias-Backend/utils"
)

type IUsuario interface {
	NovoUsuario(req *modelApresentacao.Usuario) error
	ListarUsuarios(params *utils.RequestParams) (*modelApresentacao.ListaUsuarios, error)
	//BuscarUsuario(id int) (*modelApresentacao.ReqUsuario, error)
	InativarUsuario(id int) error
	//AtualizarUsuario(id int, req *modelApresentacao.ReqUsuario) error
}
