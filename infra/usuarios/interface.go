package usuarios

import(
	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/usuarios/model"
)

type IUsuario interface {
	NovoUsuario(req *modelApresentacao.ReqUsuario) error
	ListarUsuarios(order int) ([]modelApresentacao.ReqUsuario, error)
	//BuscarUsuario(id int) (*modelApresentacao.ReqUsuario, error)
	//InativarUsuario(id int) error
	//AtualizarUsuario(id int, req *modelApresentacao.ReqUsuario) error
}