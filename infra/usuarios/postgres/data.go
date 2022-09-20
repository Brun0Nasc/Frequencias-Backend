package postgres

import (
	"database/sql"
	"fmt"

	modelData "github.com/Brun0Nasc/Frequencias-Backend/infra/usuarios/model"
	//modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/usuarios/model"
) 

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