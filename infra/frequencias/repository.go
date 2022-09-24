package frequencias

import (
	"database/sql"

	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/frequencias/model"
	modelData "github.com/Brun0Nasc/Frequencias-Backend/infra/frequencias/model"
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

func (r *repositorio) NovaFrequencia(req *modelApresentacao.Frequencia) error {
	return r.Data.NovaFrequencia(&modelData.Frequencia{
		UsuarioID: req.UsuarioID,
	})
}

func (r *repositorio) ListarFrequenciasUsuario(idUser *int64) (*modelApresentacao.ListaFrequencias, error) {
	return r.Data.ListarFrequenciasUsuario(idUser)
}
