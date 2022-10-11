package postgres

import (
	"database/sql"
	"fmt"
	"strconv"

	sq "github.com/Masterminds/squirrel"

	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/usuarios/model"
	modelData "github.com/Brun0Nasc/Frequencias-Backend/infra/usuarios/model"
	"github.com/Brun0Nasc/Frequencias-Backend/utils"
)

var usuario = modelApresentacao.Usuario{}

type DBUsuarios struct {
	DB *sql.DB
}

func (postgres *DBUsuarios) NovoUsuario(req *modelData.Usuario) error {
	sqlStatement := `INSERT INTO usuarios
	(tipo, nome, email, senha) VALUES ($1::INTEGER, $2::VARCHAR(100), 
	$3::VARCHAR(200), $4::TEXT)`

	tipo, err := strconv.Atoi(*req.Tipo)
	if err != nil {
		return err
	}

	_, err = postgres.DB.Exec(sqlStatement, tipo, req.Nome, req.Email, req.Senha)
	if err != nil {
		return err
	}

	fmt.Println("Usuario Adicionado")
	return nil
}

func (pg *DBUsuarios) ListarUsuarios(params *utils.RequestParams) (res *modelApresentacao.ListaUsuarios, err error) {
	var (
		ordem, ordenador string
	)

	if params.TemFiltro("order") {
		ordem = params.Filters["order"][0]
	}

	if params.TemFiltro("orderBy") {
		ordenador = params.Filters["orderBy"][0]
	}

	consulta := sq.
		Select("US.*").
		From("usuarios US").
		OrderBy(ordenador + " " + ordem).
		PlaceholderFormat(sq.Dollar)

	// Adicionando condições a partir dos filtros
	utils.TransformFilterInQuery(params, &consulta,
		utils.BuildFilter("removido", "(FR.removed_at IS NOT NULL)", utils.FlagFilterEq),
	)

	sqlStmt, sqlValues, err := consulta.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := pg.DB.Query(sqlStmt, sqlValues...)
	if err != nil {
		return nil, err
	}

	res = &modelApresentacao.ListaUsuarios{
		Dados: make([]modelApresentacao.Usuario, 0),
	}

	for rows.Next() {
		if err := rows.Scan(&usuario.ID, &usuario.Tipo, &usuario.Nome, &usuario.Email, &usuario.Senha, &usuario.CreatedAt, &usuario.UpdatedAt, &usuario.RemovedAt); err != nil {
			if err == sql.ErrNoRows {
				return res, nil
			}
			return nil, err
		}
		res.Dados = append(res.Dados, usuario)
	}

	fmt.Println("Listagem de usuários deu certo")
	return
}

func (postgres *DBUsuarios) InativarUsuario(id int) error {
	sqlStatement := `UPDATE usuarios SET removed_at = now() WHERE id = $1`

	_, err := postgres.DB.Exec(sqlStatement, id)
	if err != nil {
		return err
	}

	fmt.Println("Inativar usuário deu certo")
	return nil
}
