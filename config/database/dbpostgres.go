package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

/*
 *Para testes Localhost

var host string = "localhost"
var user string = "root"
var password string = "root"
var dbname string = "db_frequencia"
const port = 5432
*/

var host string = os.Getenv("DB_HOST")
var user string = os.Getenv("DB_USER")
var password string = os.Getenv("DB_PASSWORD")
var dbname string = os.Getenv("DB_NAME")
const port = 5432

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