package postgres

import (
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/frequencias/model"
	modelData "github.com/Brun0Nasc/Frequencias-Backend/infra/frequencias/model"
)

type DBFrequencias struct {
	DB *sql.DB
}

func (pg *DBFrequencias) NovaFrequencia(req *modelData.Frequencia) (err error) {
	sqlStatement, sqlValues, err := sq.
		Insert("frequencias").
		Columns("usuario_id").
		Values(req.UsuarioID).ToSql()
	if err != nil {
		return
	}

	_, err = pg.DB.Exec(sqlStatement, sqlValues...)
	if err != nil {
		return err
	}

	fmt.Println("Frequencia adicionada")
	return
}

func (pg *DBFrequencias) ListarFrequenciasUsuario(idUser *int64) (res *modelApresentacao.ListaFrequencias, err error) {
	var frequencia = modelApresentacao.Frequencia{}

	sqlStatement, sqlValues, err := sq.
		Select("FR.*").
		From("frequencias FR").
		Join("usuarios US ON US.ID = FR.usuario_id").
		Where(sq.Eq{
			"US.id": idUser,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	row, err := pg.DB.Query(sqlStatement, sqlValues...)
	if err != nil {
		return nil, err
	}

	res = &modelApresentacao.ListaFrequencias{
		Dados: make([]modelApresentacao.Frequencia, 0),
	}

	for row.Next() {
		if err := row.Scan(&frequencia.ID, &frequencia.UsuarioID, &frequencia.CreatedAt, &frequencia.Nome); err != nil {
			if err == sql.ErrNoRows {
				return res, nil
			}
			return nil, err
		}

		res.Dados = append(res.Dados, frequencia)
	}

	return
}
