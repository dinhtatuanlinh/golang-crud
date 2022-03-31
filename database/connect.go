package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func DBConn() (db *sql.DB) {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "crud"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:3306)/"+dbName)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("ok!!")
	}

	return db
}
