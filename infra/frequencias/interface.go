package frequencias

import (
	domain "github.com/Brun0Nasc/Frequencias-Backend/domain/frequencias/model"
	"github.com/Brun0Nasc/Frequencias-Backend/utils"
)

type IFrequencia interface {
	NovaFrequencia() error
	PegarFrequenciaMaisRecente() (*int, error)
	ListarFrequencias(params *utils.RequestParams) (*domain.ListaFrequencias, error)
}
