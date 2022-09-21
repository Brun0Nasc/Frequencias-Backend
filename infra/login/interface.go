package login

import (
	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/login/model"
)
type ILogin interface {
	LoginUsuario(req *modelApresentacao.Login) (*modelApresentacao.Login, error)
}