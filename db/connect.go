package db

import (
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

func Connect(host string, port string, user string, password string, dbname string) (err error) {
	var psqlInfo string = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname)

	DB, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()

	if err != nil {
		log.Fatal(err)
	}

	return
}
