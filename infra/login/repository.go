package login

import (
	"database/sql"
	modelData "github.com/Brun0Nasc/Frequencias-Backend/infra/login/model"
	"github.com/Brun0Nasc/Frequencias-Backend/infra/login/postgres"

	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/usuarios/model"
)

type repositorio struct {
	Data *postgres.DBLogin
}

func novoRepo(novoDB *sql.DB) *repositorio {
	return &repositorio{
		Data: &postgres.DBLogin{DB: novoDB},
	}
}

func (r *repositorio) LoginUsuario(req *modelApresentacao.Usuario) (*modelApresentacao.Usuario, error) {
	return r.Data.LoginUsuario(&modelData.ReqLogin{Email: req.Email, Senha: req.Senha})
}
