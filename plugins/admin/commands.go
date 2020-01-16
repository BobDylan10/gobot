package admin

import (
	"testbot/server"

	"testbot/players"

	"testbot/log"
)

func onBonjour(id int, args string) {
	server.BigText(args)
}

//TODO: Remove this from handler once a superadmin was created
func onIamgod(id int, args string) {
	//First we check that no-one has ever already done this
	pl, ok := players.GetPlayer(id)
	if (ok) {
		pls := players.GetPlayersOfLevel(100)
		if (len(pls) == 0) {
			log.Log(log.LOG_INFO, "Superadmin to be created !")
			pl.SetPlayerLevel(100)
		} else {
			log.Log(log.LOG_VERBOSE, "Superadmin already exists !")
		}
	} else {
		log.Log(log.LOG_ERROR, "Client ID does not exist when doing a lookup")
	}
	
}