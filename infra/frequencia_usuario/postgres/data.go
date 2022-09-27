package postgres

import (
	"database/sql"
	"fmt"
	"time"

	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/frequencia_usuario/model"
	modelData "github.com/Brun0Nasc/Frequencias-Backend/infra/frequencia_usuario/model"
	"github.com/Brun0Nasc/Frequencias-Backend/utils"
	sq "github.com/Masterminds/squirrel"
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

func Filtro(consulta *sq.SelectBuilder, params *utils.RequestParams){
	var data string
	if params.TemFiltro("data") {
		data = params.Filters["data"][0]
	}
	consulta.Where(sq.Eq{
		"(FR.data_atual)": data,
	})
}

func (pg *DBFrequenciaUsuario) ListaFrequenciasData(params *utils.RequestParams) (res *modelApresentacao.ListaUsuarioFrequencia, err error) {
	var (
		ordem, ordenador string
	)

	if params.TemFiltro("order") {
		ordem = params.Filters["order"][0]
	}

	if params.TemFiltro("orderBy") {
		ordenador = params.Filters["orderBy"][0]
	}

	

	sqlStmt, sqlValues, err := sq.
		Select("FU.id, FU.usuario_id, US.nome, FU.frequencia_id, FR.data_atual data_frequencia, FU.entrada, FU.saida").
		From("frequencia_usuario FU").
		Join("usuarios US ON US.id=FU.usuario_id").
		Join("frequencias FR ON FR.id=FU.frequencia_id").
		OrderBy(ordenador + " " + ordem).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return
	}

	rows, err := pg.DB.Query(sqlStmt, sqlValues...)
	if err != nil {
		return
	}

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
