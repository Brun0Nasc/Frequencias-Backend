package postgres

import (
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/usuarios/model"
	modelData "github.com/Brun0Nasc/Frequencias-Backend/infra/usuarios/model"
)

var usuario = modelApresentacao.Usuario{}

type DBUsuarios struct {
	DB *sql.DB
}

func (postgres *DBUsuarios) NovoUsuario(req *modelData.Usuario) error {
	sqlStatement := `INSERT INTO usuarios
	(tipo, nome, email, senha) VALUES ($1::INTEGER, $2::VARCHAR(100), 
	$3::VARCHAR(200), $4::TEXT)`

	_, err := postgres.DB.Exec(sqlStatement, req.Tipo, req.Nome, req.Email, req.Senha)
	if err != nil {
		return err
	}

	fmt.Println("Usuario Adicionado")
	return nil
}

func (pg *DBUsuarios) ListarUsuarios(order int) (res *modelApresentacao.ListaUsuarios, err error) {
	var sqlStmt string
	var sqlValues []interface{}

	switch order {
	case 1:
		sqlStmt, sqlValues, err = sq.
		Select("US.*").
		From("usuarios US").
		Where(sq.NotEq{
			"US.tipo":0,
		}).
		OrderBy("UPPER(US.nome)").
		PlaceholderFormat(sq.Dollar).
		ToSql()
		//sqlStatement = `SELECT * FROM usuarios WHERE tipo != 0 ORDER BY UPPER(nome);` // listar usuários ativos em ordem A_Z
	case 2:
		sqlStmt, sqlValues, err = sq.
		Select("US.*").
		From("usuarios US").
		Where(sq.NotEq{
			"US.tipo":0,
		}).
		OrderBy("UPPER(US.nome) DESC").
		PlaceholderFormat(sq.Dollar).
		ToSql()
		//sqlStatement = `SELECT * FROM usuarios WHERE tipo != 0 ORDER BY UPPER(nome) DESC;` // listar usuários ativos em ordem Z_A
	case 3:
		sqlStmt, sqlValues, err = sq.
		Select("US.*").
		From("usuarios US").
		Where(sq.Eq{
			"US.tipo":0,
		}).
		OrderBy("UPPER(US.nome)").
		PlaceholderFormat(sq.Dollar).
		ToSql()
		//sqlStatement = `SELECT * FROM usuarios WHERE tipo = 0 ORDER BY UPPER(nome);` // listar usuários inativos em ordem A_Z
	case 4:
		sqlStmt, sqlValues, err = sq.
		Select("US.*").
		From("usuarios US").
		Where(sq.Eq{
			"US.tipo":0,
		}).
		OrderBy("UPPER(US.nome) DESC").
		PlaceholderFormat(sq.Dollar).
		ToSql()
		//sqlStatement = `SELECT * FROM usuarios WHERE tipo = 0 ORDER BY UPPER(nome) DESC;` // listar usuários inativos em ordem Z_A
	case 5:
		sqlStmt, sqlValues, err = sq.
		Select("US.*").
		From("usuarios US").
		Where(sq.NotEq{
			"US.tipo":0,
		}).
		OrderBy("US.created_at").
		PlaceholderFormat(sq.Dollar).
		ToSql()
		//sqlStatement = `SELECT * FROM usuarios WHERE tipo != 0 ORDER BY created_at;` // listar usuários ativos em orderm de criação crescente
	case 6:
		sqlStmt, sqlValues, err = sq.
		Select("US.*").
		From("usuarios US").
		Where(sq.NotEq{
			"US.tipo":0,
		}).
		OrderBy("US.created_at DESC").
		PlaceholderFormat(sq.Dollar).
		ToSql()
		//sqlStatement = `SELECT * FROM usuarios WHERE tipo != 0 ORDER BY created_at DESC;` // listar usuários ativos em orderm de criação decrescente
	case 7:
		sqlStmt, sqlValues, err = sq.
		Select("US.*").
		From("usuarios US").
		Where(sq.Eq{
			"US.tipo":0,
		}).
		OrderBy("US.created_at").
		PlaceholderFormat(sq.Dollar).
		ToSql()
		//sqlStatement = `SELECT * FROM usuarios WHERE tipo = 0 ORDER BY created_at;` // listar usuários inativos em orderm de criação crescente
	case 8:
		sqlStmt, sqlValues, err = sq.
		Select("US.*").
		From("usuarios US").
		Where(sq.Eq{
			"US.tipo":0,
		}).
		OrderBy("US.created_at DESC").
		PlaceholderFormat(sq.Dollar).
		ToSql()
		//sqlStatement = `SELECT * FROM usuarios WHERE tipo = 0 ORDER BY created_at DESC;` // listar usuários inativos em orderm de criação decrescente
	default:
		return nil, fmt.Errorf("Order " + fmt.Sprint(order) + " inexistente.")
	}

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
		if err := rows.Scan(&usuario.ID, &usuario.Tipo, &usuario.Nome, &usuario.Email, &usuario.Senha, &usuario.CreatedAt, &usuario.UpdatedAt); err != nil {
			if err == sql.ErrNoRows{
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
	sqlStatement := `UPDATE usuarios SET tipo = 0 WHERE id = $1`

	_, err := postgres.DB.Exec(sqlStatement, id)
	if err != nil {
		return err
	}

	fmt.Println("Inativar usuário deu certo")
	return nil
}
