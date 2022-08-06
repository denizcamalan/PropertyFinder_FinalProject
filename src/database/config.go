package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Get_db() (db *sql.DB, err error) {

	dataBaseSystem := "mysql"
	dbName := "pf_project"
	dbUser := "root"
	dbPass := "95987168"
	db, err = sql.Open(dataBaseSystem, dbUser+":"+dbPass+"@tcp(127.0.0.1:8889)/"+dbName)

	// db, err = sql.Open("mysql","root:95987168@tcp(127.0.0.1:8889)/pf_project")
	// db.Ping()
	
	return 
}


