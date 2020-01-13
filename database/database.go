package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testbot/log"
)

var Database *sql.DB

// Sets the wanted database
func SetDatabase(dsn string) {
	log.Log(log.LOG_INFO, "Setting up database")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Log(log.LOG_ERROR, "Cannot open database")
	} //TODO: Need to do better
	err = db.Ping()
	if err != nil {
		log.Log(log.LOG_ERROR, "Cannot open database")
	}
	Database = db
}

func CloseDatabase() {
	Database.Close()
}