package frequencias

import (
	"database/sql"
) 

func NovoRepo(DB *sql.DB) IFrequencia {
	return novoRepo(DB)
}