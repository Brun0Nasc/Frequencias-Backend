package postgres

import (
	"database/sql"
	"fmt"

	domain "github.com/Brun0Nasc/Frequencias-Backend/domain/frequencias/model"
	"github.com/Brun0Nasc/Frequencias-Backend/utils"
	sq "github.com/Masterminds/squirrel"
)

type DBFrequencias struct {
	DB *sql.DB
}

func (pg *DBFrequencias) NovaFrequencia() (err error) {
	sqlStatement, _, err := sq.
		Insert("frequencias").
		Columns("data_atual").
		Values("current_date").ToSql()
	if err != nil {
		return
	}

	_, err = pg.DB.Exec(sqlStatement)
	if err != nil {
		return err
	}

	fmt.Println("Frequencia adicionada")
	return
}

func (pg *DBFrequencias) PegarFrequenciaMaisRecente() (res *int, err error) {
	sqlStatement, _, err := sq.
		Select("FR.id").
		From("frequencias FR").
		OrderBy("FR.id desc").
		Limit(1).
		ToSql()

	if err != nil {
		fmt.Println(err)
	}

	row := pg.DB.QueryRow(sqlStatement)

	if err := row.Scan(&res); err != nil { // ! Erro: destination pointer is nil
		if err == sql.ErrNoRows {
			return res, nil
		}
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Aqui Funciona 3 ")
	return
}

func (pg *DBFrequencias) ListarFrequencias(params *utils.RequestParams) (res *domain.ListaFrequencias, err error) {
	var (
		conditions *sq.And
		frequencia = domain.Frequencia{}
	)
	if params.TemFiltro("existe_frequencia_hoje") {
		*conditions = sq.And{
			sq.Eq{
				"FR.data_atual": sq.Expr("NOW()"),
			},
		}
	}

	sqlStatement, _, err := sq.
		Select("FR.id").
		From("frequencias FR").
		Where(conditions).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return
	}

	rows, err := pg.DB.Query(sqlStatement)

	for rows.Next() {
		if err = rows.Scan(&frequencia.ID); err != nil {
			if err == sql.ErrNoRows {
				return res, nil
			}
			return
		}

		res.Dados = append(res.Dados, frequencia)
	}

	return
}
