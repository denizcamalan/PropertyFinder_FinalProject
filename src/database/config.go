package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Get_db() (db *sql.DB, err error) {

	// dataBaseSystem := "mysql"
	// dbName := "database"
	// dbUser := "user"
	// dbPass := "password"
	// db, err = sql.Open(dataBaseSystem, dbUser+":"+dbPass+"@tcp(127.0.0.1:8095)/"+dbName)

	db, err = sql.Open("mysql","user:password@tcp(127.0.0.1:8095)/pf_project")
	db.Ping()
	
	return 
}


