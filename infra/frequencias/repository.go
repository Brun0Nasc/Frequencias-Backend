package frequencias

import (
	"database/sql"

	domain "github.com/Brun0Nasc/Frequencias-Backend/domain/frequencias/model"
	"github.com/Brun0Nasc/Frequencias-Backend/infra/frequencias/postgres"
	"github.com/Brun0Nasc/Frequencias-Backend/utils"
)

type repositorio struct {
	Data *postgres.DBFrequencias
}

func novoRepo(novoDB *sql.DB) *repositorio {
	return &repositorio{
		Data: &postgres.DBFrequencias{DB: novoDB},
	}
}

func (r *repositorio) NovaFrequencia() error {
	return r.Data.NovaFrequencia()
}

func (r *repositorio) PegarFrequenciaMaisRecente() (*int, error) {
	return r.Data.PegarFrequenciaMaisRecente()
}

func (r *repositorio) ListarFrequencias(params *utils.RequestParams) (*domain.ListaFrequencias, error) {
	return r.Data.ListarFrequencias(params)
}
