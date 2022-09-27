package postgres

import (
	"database/sql"
	"fmt"

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
