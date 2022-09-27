package postgres

import (
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	modelData "github.com/Brun0Nasc/Frequencias-Backend/infra/frequencia_usuario/model"
)

type DBFrequenciaUsuario struct {
	DB *sql.DB
}

func (pg *DBFrequenciaUsuario) NovaFrequenciaUsuario(req *modelData.FrequenciaUsuario) (err error) {
	sqlStatement, sqlValues, err := sq.
		Insert("frequencia_usuario").
		Columns("usuario_id, frequencia_id").
		Values(req.UsuarioID, req.FrequenciaID).
		PlaceholderFormat(sq.Dollar).
		ToSql()
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

/*func (pg *DBFrequencias) ListarFrequenciasUsuario(idUser *int64) (res *modelApresentacao.ListaFrequencias, err error) {
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
*/
