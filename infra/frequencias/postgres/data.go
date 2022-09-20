package postgres

import (
	"database/sql"
	"fmt"

	modelData "github.com/Brun0Nasc/Frequencias-Backend/infra/frequencias/model"
)

type DBFrequencias struct {
	DB *sql.DB
}

func (postgres *DBFrequencias) NovaFrequencia(req *modelData.Frequencia) error {
	sqlStatement := `INSERT INTO frequencias 
					(usuario_id)
					VALUES
					($1::INTEGER)`

	_, err := postgres.DB.Exec(sqlStatement, req.UsuarioID)
	if err != nil {
		return err
	}

	fmt.Println("Frequencia adicionada")
	return nil
}