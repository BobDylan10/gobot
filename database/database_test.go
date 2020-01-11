package database

import (
	"testing"
	"fmt"
)

func TestDatabase(t *testing.T) {
	SetDatabase("gobot:gobot@/gobot_db")
	stmtOut, err := Database.Prepare("SELECT player_id, alias, last_connection, connections, guid FROM players WHERE guid = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	var player_id int
	var alias string
	var date string
	var connections int
	var guid string

	err = stmtOut.QueryRow("MYGUID").Scan(&player_id, &alias, &date, &connections, &guid) // WHERE number = 13
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("alias is: %d, %s, %s, %d, %s\n", player_id, alias, date, connections, guid)

	CloseDatabase()
}