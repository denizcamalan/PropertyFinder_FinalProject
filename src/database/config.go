package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Get_db() (db *sql.DB, err error) {

	// dataBaseSystem := "mysql"
	// dbName := "pf_project"
	// dbUser := "root"
	// dbPass := "root"
	// db, err = sql.Open(dataBaseSystem, dbUser+":"+dbPass+"@tcp(172.22.0.1:8889)/"+dbName)

	db, err = sql.Open("mysql","root:root@tcp(172.22.0.2:8888)/pf_project")
	if err != nil {
		return db,err
	}
	//db.SetMaxIdleConns(-1)
	//db.SetConnMaxLifetime(time.Second)
	db.Ping()	

	return  db, nil
}