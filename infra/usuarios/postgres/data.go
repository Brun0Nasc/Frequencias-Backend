package postgres

import (
	"database/sql"
	"fmt"

	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/usuarios/model"
	modelData "github.com/Brun0Nasc/Frequencias-Backend/infra/usuarios/model"
)

var usuario = modelApresentacao.ReqUsuario{}

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

func (postgres *DBUsuarios) ListarUsuarios(order int) ([]modelApresentacao.ReqUsuario, error) {
	var sqlStatement string

	switch order {
	case 1:
		sqlStatement = `SELECT * FROM usuarios WHERE tipo != 0 ORDER BY UPPER(nome);` // listar usuários ativos em ordem A_Z
	case 2:
		sqlStatement = `SELECT * FROM usuarios WHERE tipo != 0 ORDER BY UPPER(nome) DESC;` // listar usuários ativos em ordem Z_A
	case 3:
		sqlStatement = `SELECT * FROM usuarios WHERE tipo = 0 ORDER BY UPPER(nome);` // listar usuários inativos em ordem A_Z
	case 4:
		sqlStatement = `SELECT * FROM usuarios WHERE tipo = 0 ORDER BY UPPER(nome) DESC;` // listar usuários inativos em ordem Z_A
	case 5:
		sqlStatement = `SELECT * FROM usuarios WHERE tipo != 0 ORDER BY created_at;` // listar usuários ativos em orderm de criação crescente
	case 6:
		sqlStatement = `SELECT * FROM usuarios WHERE tipo != 0 ORDER BY created_at DESC;` // listar usuários ativos em orderm de criação decrescente
	case 7:
		sqlStatement = `SELECT * FROM usuarios WHERE tipo = 0 ORDER BY created_at;` // listar usuários inativos em orderm de criação crescente
	case 8:
		sqlStatement = `SELECT * FROM usuarios WHERE tipo = 0 ORDER BY created_at DESC;` // listar usuários inativos em orderm de criação decrescente
	default:
		return nil, fmt.Errorf("Order " + fmt.Sprint(order) + " inexistente.")
	}

	rows, err := postgres.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	var res = []modelApresentacao.ReqUsuario{}

	for rows.Next() {
		if err := rows.Scan(&usuario.ID, &usuario.Tipo, &usuario.Nome, &usuario.Email, &usuario.Senha, &usuario.CreatedAt, &usuario.UpdatedAt); err != nil {
			return nil, err
		}
		res = append(res, usuario)
	}
	fmt.Println("Listagem de usuários deu certo")
	return res, nil
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
