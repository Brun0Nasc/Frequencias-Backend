package frequencia_usuario

import(
	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/frequencia_usuario/model"
	"github.com/Brun0Nasc/Frequencias-Backend/utils"
)

type IFrequenciaUsuario interface{
	NovaFrequenciaUsuario(req *modelApresentacao.ReqFrequenciaUsuario) error
	ListaFrequenciasData(params *utils.RequestParams) (*modelApresentacao.ListaUsuarioFrequencia, error)
}