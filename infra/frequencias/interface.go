package frequencias

import (
	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/frequencias/model"
)

type IFrequencia interface {
	NovaFrequencia(req *modelApresentacao.Frequencia) error
	ListarFrequenciasUsuario(idUser *int64) (*modelApresentacao.ListaFrequencias, error)
}
