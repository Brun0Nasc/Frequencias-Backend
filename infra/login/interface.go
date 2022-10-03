package login

import (
	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/usuarios/model"
)
type ILogin interface {
	LoginUsuario(req *modelApresentacao.Usuario) (*modelApresentacao.Usuario, error)
}