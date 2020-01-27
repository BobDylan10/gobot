
package stats

import (
	"testbot/log"
	"testbot/database"
)

func initTable() {
	log.Log(log.LOG_INFO, "Initializing stats table")
	stmtInit := database.Prepare(`CREATE TABLE IF NOT EXISTS connectiontimes (
		id INT AUTO_INCREMENT PRIMARY KEY,
		player_id INT NOT NULL,
		c_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		minutes FLOAT NOT NULL,
		FOREIGN KEY (player_id)
		   REFERENCES players(player_id)
	 ) ENGINE=InnoDB;`)
	defer stmtInit.Close()
	_, err := stmtInit.Exec()
	if err != nil {
		log.Log(log.LOG_ERROR, "Error initializing stats table")
	}
}

func newTimeSpan(did int, minutes float64) {
	log.Log(log.LOG_INFO, "New connection for did:", did, "for", minutes, "minutes")
	stmtIns := database.Prepare("INSERT INTO connectiontimes(player_id, minutes) VALUES( ?, ? )") // ? = placeholder
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates


	_, err := stmtIns.Exec() // Insert tuples (i, i^2)
		if err != nil {
			log.Log(log.LOG_ERROR, "Error creating connection for", did, minutes)
		}
}