package players

import (
	"database/sql"
	"testbot/database"
	"fmt"
)

func getPlayerDid(guid string, name string) {
	stmtOut, err := database.Database.Prepare("SELECT player_id, alias, last_connection, connections FROM players WHERE guid = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	var did int
	var alias string
	var date string
	var connections int

	err = stmtOut.QueryRow(guid).Scan(&did, &alias, &date, &connections) // WHERE number = 13
	if err != nil {
		if (err == sql.ErrNoRows) {
			fmt.Println("Creating player with GUID", guid)
			createPlayer(guid, name)
		} else {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	} else {
		fmt.Println("Found player !")
	}
	err = stmtOut.QueryRow(guid).Scan(&did, &alias, &date, &connections) // WHERE number = 13
	if err != nil { //Here it should exist in all cases
			panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("alias is: %d, %s, %s, %d\n", did, alias, date, connections)
}

func createPlayer(guid string, name string) {
	stmtIns, err := database.Database.Prepare("INSERT INTO players(alias, guid) VALUES( ?, ? )") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates


	_, err = stmtIns.Exec(name, guid) // Insert tuples (i, i^2)
		if err != nil {
			fmt.Println("Error creating player", name, guid)
		}
}