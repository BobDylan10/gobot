package players

import (
	"database/sql"
	"testbot/database"
	"testbot/log"
	"time"
)

//This function will always succeed, because it creates the player if it's not already on the database
func getPlayer(guid string, name string, attributes map[string]string, player *Player) {
	stmtOut := database.Prepare("SELECT player_id, alias, level, last_connection, connections FROM players WHERE guid = ?")
	defer stmtOut.Close()

	var did int
	var alias string
	var level int
	var date time.Time
	var connections int

	err := stmtOut.QueryRow(guid).Scan(&did, &alias, &level, &date, &connections) // WHERE number = 13
	if err != nil {
		if err == sql.ErrNoRows {
			createPlayer(guid, name)
		} else {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	} else {
		log.Log(log.LOG_DEBUG, "Found player with guid", guid)
	}
	err = stmtOut.QueryRow(guid).Scan(&did, &alias, &level, &date, &connections) // WHERE number = 13
	if err != nil {                                                              //Here it should exist in all cases
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	player.did = did
	player.level = level
	player.connections = connections
	player.isBot = false
	player.lastConnection = date
	player.name = name
	player.guid = guid
	player.toBeDeleted = false
	player.attributes = attributes
}

func createPlayer(guid string, name string) {
	log.Log(log.LOG_INFO, "Creating new player", name, "with guid", guid)
	stmtIns := database.Prepare("INSERT INTO players(alias, guid, first_seen) VALUES( ?, ?, ? )") // ? = placeholder
	defer stmtIns.Close()                                                                         // Close the statement when we leave main() / the program terminates

	_, err := stmtIns.Exec(name, guid, time.Now()) // Insert tuples (i, i^2)
	if err != nil {
		log.Log(log.LOG_ERROR, "Error creating player", name, guid)
	}
}

func (pl *Player) newConnection() {
	log.Log(log.LOG_DEBUG, "Updating player with did", pl.did)
	pl.connections += 1
	stmtUp := database.Prepare(
		`UPDATE players
		SET
			connections = ?,
			last_connection = ?
		WHERE player_id = ?;`) // ? = placeholder
	defer stmtUp.Close()
	_, err := stmtUp.Exec(pl.connections, time.Now(), pl.did)
	if err != nil {
		log.Log(log.LOG_ERROR, "Error update player with did", pl.did)
	}
}

func initTable() {
	log.Log(log.LOG_INFO, "Initializing players table")
	stmtInit := database.Prepare(`CREATE TABLE IF NOT EXISTS players (
		player_id INT AUTO_INCREMENT PRIMARY KEY,
		alias VARCHAR(32) NOT NULL,
		level INT NOT NULL DEFAULT 1,
		first_seen TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		last_connection TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		connections INT NOT NULL DEFAULT 0,
		guid VARCHAR(36) NOT NULL UNIQUE
	 ) ENGINE=InnoDB;`)
	defer stmtInit.Close()
	_, err := stmtInit.Exec()
	if err != nil {
		log.Log(log.LOG_ERROR, "Error initializing player table")
	}
}

func GetPlayersOfLevel(lvl int) []Player {
	res := []Player{}
	stmtGetLvl := database.Prepare("SELECT player_id, alias, guid FROM players WHERE level = ? ORDER BY RAND() LIMIT 10")
	defer stmtGetLvl.Close()

	rows, err := stmtGetLvl.Query(lvl)
	if err != nil {
		if err == sql.ErrNoRows {
			return res
		} else {
			log.Log(log.LOG_ERROR, "Error querying players by level")
		}
	}
	for rows.Next() {
		var pl Player
		rows.Scan(&pl.did, &pl.name, &pl.guid)
		res = append(res, pl)
	}
	return res
}

func (pl *Player) SetPlayerLevel(lvl int) {
	log.Log(log.LOG_INFO, "Updating player level with did", pl.did, "to", lvl)
	pl.level = lvl
	stmtUp := database.Prepare(
		`UPDATE players
		SET
			level = ?
		WHERE player_id = ?;`) // ? = placeholder
	defer stmtUp.Close()
	_, err := stmtUp.Exec(lvl, pl.did)
	if err != nil {
		log.Log(log.LOG_ERROR, "Error update player with new level", pl.did)
	}
}

func GetConnectedPlayers() ([]int, []Player) {
	indices := make([]int, len(players))
	ret := make([]Player, len(players))
	c := 0
	for i, pl := range players {
		ret[c] = *pl
		indices[c] = i
		c++
	}
	return indices, ret
}
