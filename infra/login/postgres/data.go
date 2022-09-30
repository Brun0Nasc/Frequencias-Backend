package postgres

import (
	"database/sql"
	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/usuarios/model"
	modelData "github.com/Brun0Nasc/Frequencias-Backend/infra/login/model"
	sq "github.com/Masterminds/squirrel"
)

type DBLogin struct {
	DB *sql.DB
}

func (postgres *DBLogin) LoginUsuario(req *modelData.ReqLogin) (*modelApresentacao.Usuario, error) {
	sqlStatement, sqlValues, err := sq.
		Select("US.id, US.tipo, US.nome, US.email, US.senha").
		From("usuarios US").
		Where(sq.Eq{
			"email": req.Email,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	var res = &modelApresentacao.Usuario{}
	row := postgres.DB.QueryRow(sqlStatement, sqlValues...)

	if err := row.Scan(&res.ID, &res.Tipo, &res.Nome, &res.Email, &res.Senha); err != nil {
		return nil, err
	}

	if res.Senha != req.Senha {
		return nil, nil
	}

	return res, nil
}
