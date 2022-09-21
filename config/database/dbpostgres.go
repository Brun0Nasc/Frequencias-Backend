package database

import (
	"database/sql"
	"fmt"
	"log"
	//"os"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "db_ponto"
)

func Conectar() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " +
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	return db
}