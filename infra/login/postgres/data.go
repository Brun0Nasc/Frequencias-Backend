package postgres

import (
	"database/sql"
	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/login/model"
	modelData "github.com/Brun0Nasc/Frequencias-Backend/infra/login/model"
)

type DBLogin struct {
	DB *sql.DB
}

func (postgres *DBLogin) LoginUsuario(req *modelData.ReqLogin) (*modelApresentacao.Login, error) {
	sqlStatement := `SELECT * FROM usuarios WHERE email = $1::VARCHAR(80);`
	user := modelApresentacao.Login{}
	row := postgres.DB.QueryRow(sqlStatement, req.Email)

	if err := row.Scan(&user.ID, &user.Tipo, &user.Nome, &user.Email, &user.Senha, &user.CreatedAt, &user.UpdatedAt, &user.RemovedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			return nil, err
		}
	}

	if user.Senha != req.Senha {
		return nil, nil
	}

	return &user, nil
}