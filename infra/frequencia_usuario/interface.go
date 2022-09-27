package frequencia_usuario

import(
	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/frequencia_usuario/model"
)

type IFrequenciaUsuario interface{
	NovaFrequenciaUsuario(req *modelApresentacao.ReqFrequenciaUsuario) error
}