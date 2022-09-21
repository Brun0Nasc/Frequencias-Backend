package usuarios

import (
	"database/sql"
	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/usuarios/model"
	modelData "github.com/Brun0Nasc/Frequencias-Backend/infra/usuarios/model"
	"github.com/Brun0Nasc/Frequencias-Backend/infra/usuarios/postgres"
)

type repositorio struct {
	Data *postgres.DBUsuarios
}

func novoRepo(novoDB *sql.DB) *repositorio {
	return &repositorio{
		Data: &postgres.DBUsuarios{DB: novoDB},
	}
}

func (r *repositorio) NovoUsuario(req *modelApresentacao.ReqUsuario) error {
	return r.Data.NovoUsuario(&modelData.Usuario{
		Tipo:  req.Tipo,
		Nome:  req.Nome,
		Email: req.Email,
		Senha: &req.Senha,
	})
}

func (r *repositorio) ListarUsuarios(order int) ([]modelApresentacao.ReqUsuario, error) {
	return r.Data.ListarUsuarios(order)
}
