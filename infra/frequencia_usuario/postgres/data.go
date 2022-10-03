package postgres

import (
	"database/sql"
	"fmt"

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

func (pg *DBFrequenciaUsuario) ListaFrequenciasData(params *utils.RequestParams) (res *modelApresentacao.ListaUsuarioFrequencia, err error) {
	var (
		ordem, ordenador string
	)
	var frequenciaUsuario = modelApresentacao.ResFrequenciaUsuario{}

	if params.TemFiltro("order") {
		ordem = params.Filters["order"][0]
	}

	if params.TemFiltro("orderBy") {
		ordenador = params.Filters["orderBy"][0]
	}

	consulta := sq.
		Select(
			"FU.id",
			"FU.usuario_id",
			"US.nome",
			"FU.frequencia_id",
			"FR.data_atual",
			"FU.entrada",
			"FU.saida",
		).From("frequencia_usuario FU").
		Join("usuarios US ON US.id=FU.usuario_id").
		Join("frequencias FR ON FR.id=FU.frequencia_id").
		OrderBy(ordenador + " " + ordem).
		PlaceholderFormat(sq.Dollar)

	// Adicionando condições a partir dos filtros
	utils.TransformFilterInQuery(params, &consulta,
		utils.BuildFilter("data", "FR.data_atual::DATE", utils.FlagFilterEq),
	)

	sqlStmt, sqlValues, err := consulta.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := pg.DB.Query(sqlStmt, sqlValues...)
	if err != nil {
		return nil, err
	}

	res = &modelApresentacao.ListaUsuarioFrequencia{
		Dados: make([]modelApresentacao.ResFrequenciaUsuario, 0),
	}
	for rows.Next() {
		if err = rows.Scan(&frequenciaUsuario.ID, &frequenciaUsuario.UsuarioID, &frequenciaUsuario.Nome, &frequenciaUsuario.FrequenciaID,
			&frequenciaUsuario.DataFrequencia, &frequenciaUsuario.Entrada, &frequenciaUsuario.Saida); err != nil {
			if err == sql.ErrNoRows {
				return res, nil
			}
			return nil, err
		}
		res.Dados = append(res.Dados, frequenciaUsuario)
	}

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
