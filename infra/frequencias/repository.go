package frequencias

import (
	"database/sql"

	"github.com/Brun0Nasc/Frequencias-Backend/infra/frequencias/postgres"
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

func (r *repositorio) PegarFrequenciaMaisRecente() (*int, error){
	return r.Data.PegarFrequenciaMaisRecente()
}
