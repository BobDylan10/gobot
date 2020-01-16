package players

import (
	"database/sql"
	"testbot/database"
	"time"
	"testbot/log"
)

func getPlayer(guid string, name string, player *player){
	stmtOut := database.Prepare("SELECT player_id, alias, level, last_connection, connections FROM players WHERE guid = ?")
	defer stmtOut.Close()

	var did int
	var alias string
	var level int
	var date time.Time
	var connections int

	err := stmtOut.QueryRow(guid).Scan(&did, &alias, &level, &date, &connections) // WHERE number = 13
	if err != nil {
		if (err == sql.ErrNoRows) {
			log.Log(log.LOG_VERBOSE, "Creating player with GUID", guid)
			createPlayer(guid, name)
		} else {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	} else {
		log.Log(log.LOG_DEBUG, "Found player with guid", guid)
	}
	err = stmtOut.QueryRow(guid).Scan(&did, &alias, &level, &date, &connections) // WHERE number = 13
	if err != nil { //Here it should exist in all cases
			panic(err.Error()) // proper error handling instead of panic in your app
	}
	player.did = did
	player.level = level
	player.connections = connections
	player.isBot = false
	player.lastConnection = date
	player.name = name
}

func createPlayer(guid string, name string) {
	stmtIns := database.Prepare("INSERT INTO players(alias, guid) VALUES( ?, ? )") // ? = placeholder
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates


	_, err := stmtIns.Exec(name, guid) // Insert tuples (i, i^2)
		if err != nil {
			log.Log(log.LOG_ERROR, "Error creating player", name, guid)
		}
}