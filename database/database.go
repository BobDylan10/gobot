package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var Database *sql.DB

// Sets the wanted database
func SetDatabase(dsn string) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Cannot open database")
	} //TODO: Need to do better
	err = db.Ping()
	if err != nil {
		fmt.Println("Did not work")
	}
	Database = db
}

func CloseDatabase() {
	Database.Close()
}