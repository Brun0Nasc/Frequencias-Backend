package postgres

import (
	"database/sql"
	"fmt"

	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/frequencias/model"
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

func (postgres *DBFrequencias) ListarFrequenciasUsuario(idUser int) ([]modelApresentacao.ReqFrequencia, error) {
	sqlStatement := `SELECT fr.*, us.nome FROM frequencias fr
	INNER JOIN usuarios us
	ON fr.usuario_id = us.id 
	WHERE fr.usuario_id = $1`

	var frequencia = modelApresentacao.ReqFrequencia{}
	var res = []modelApresentacao.ReqFrequencia{}

	row, err := postgres.DB.Query(sqlStatement, idUser)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		if err := row.Scan(&frequencia.ID, &frequencia.UsuarioID, &frequencia.CreatedAt, &frequencia.Nome); err != nil {
			return nil, err
		}
		res = append(res, frequencia)
	}
	return res, nil
}