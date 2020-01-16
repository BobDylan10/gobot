package admin

import (

	"testbot/plugins/commands"
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
		log.Log(log.LOG_INFO, "Superadmin created !")
		pl.SetPlayerLevel(100)
		commands.DeleteCommand("iamgod")
	} else {
		log.Log(log.LOG_ERROR, "Client ID does not exist when doing a lookup")
	}
}

func onSay(id int, args string) {
	server.Say("This was tested")
}