package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Get_db() (db *sql.DB, err error) {

	db, err = sql.Open("mysql","root:password@tcp(localhost:3306)/pf_project")
	if err != nil {
		log.Println(err)
	}

	return 
}