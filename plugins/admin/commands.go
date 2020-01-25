package admin

import (
	"strconv"

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

func playerToString(ind int, name string, level int) string{
	ret := strconv.FormatInt(int64(ind), 10)
	ret += ": "
	ret += name
	ret += " ["
	ret += strconv.FormatInt(int64(level), 10)
	ret += "]"
	return ret
}

func onStatus(id int, args string) {
	server.Say("Players connected:")
	indices, pls := players.GetConnectedPlayers()
	if (len(indices) != len(pls)) {
		log.Log(log.LOG_FATAL, "Players and indices don't have the same size !")
	}
	for i := 0; i < len(indices); i++ {
		server.Say(playerToString(indices[i], pls[i].GetPlayerName(), pls[i].GetPlayerLevel()))
	}
}

func onSay(id int, args string) {
	server.Say("This was tested")
}