package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testbot/log"
)

var database *sql.DB

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
	database = db
}

func Prepare(query string) (*sql.Stmt) {
	s, e := database.Prepare(query)
	if (e != nil) {
		log.Log(log.LOG_ERROR, "Error prepary query" + query)
	}
	return s
}

func CloseDatabase() {
	database.Close()
}