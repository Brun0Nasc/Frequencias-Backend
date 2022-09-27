package frequencia_usuario

import (
	"database/sql"
) 

func NovoRepo(DB *sql.DB) IFrequenciaUsuario {
	return novoRepo(DB)
}