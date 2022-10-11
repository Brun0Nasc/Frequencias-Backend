package postgres

import (
	"database/sql"
	"fmt"
	"time"

	domain "github.com/Brun0Nasc/Frequencias-Backend/domain/frequencias/model"
	"github.com/Brun0Nasc/Frequencias-Backend/utils"
	sq "github.com/Masterminds/squirrel"
)

type DBFrequencias struct {
	DB *sql.DB
}

func (pg *DBFrequencias) NovaFrequencia() (err error) {
	sqlStatement, sqlValues, err := sq.
		Insert("frequencias").
		Columns("data_atual").
		Values("NOW()").
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
		conditions = sq.And{}
		frequencia = domain.Frequencia{}
	)

	if params.TemFiltro("existe_frequencia_hoje") {
		conditions = sq.And{
			sq.Eq{
				"FR.data_atual::DATE": time.Now().Format("2006-01-02"),
			},
		}
	}

	sqlStatement, sqlValues, err := sq.
		Select("FR.id").
		From("frequencias FR").
		Where(conditions).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return
	}

	rows, err := pg.DB.Query(sqlStatement, sqlValues...)
	if err != nil {
		return
	}

	res = &domain.ListaFrequencias{Dados: make([]domain.Frequencia, 0)}
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
