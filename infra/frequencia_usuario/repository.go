package frequencia_usuario

import (
	"database/sql"

	"github.com/Brun0Nasc/Frequencias-Backend/infra/frequencia_usuario/postgres"
	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/frequencia_usuario/model"
	modelData "github.com/Brun0Nasc/Frequencias-Backend/infra/frequencia_usuario/model"
)

type repositorio struct {
	Data *postgres.DBFrequenciaUsuario
}

func novoRepo(novoDB *sql.DB) *repositorio {
	return &repositorio{
		Data: &postgres.DBFrequenciaUsuario{DB: novoDB},
	}
}

func (r *repositorio) NovaFrequenciaUsuario(req *modelApresentacao.ReqFrequenciaUsuario) error {
	return r.Data.NovaFrequenciaUsuario(&modelData.FrequenciaUsuario{
		UsuarioID: req.UsuarioID,
		FrequenciaID: req.FrequenciaID,
	})
}
