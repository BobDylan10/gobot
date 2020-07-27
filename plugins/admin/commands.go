package admin

import (
	"strconv"

	"testbot/plugins/commands"
	"testbot/server"

	"testbot/players"

	"testbot/log"
)

func onBonjour(emitter *players.Player, args string) {
	server.BigText(args)
}

func onIamgod(emitter *players.Player, args string) {
	//First we check that no-one has ever already done this
	log.Log(log.LOG_INFO, "Superadmin created !")
	emitter.SetPlayerLevel(100)
	commands.DeleteCommand("iamgod")
}

func playerToString(ind int, name string, level int) string {
	ret := strconv.FormatInt(int64(ind), 10)
	ret += ": "
	ret += name
	ret += " ["
	ret += strconv.FormatInt(int64(level), 10)
	ret += "]"
	return ret
}

func onStatus(emitter *players.Player, args string) {
	server.Say("Players connected:")
	indices, pls := players.GetConnectedPlayers()
	if len(indices) != len(pls) {
		log.Log(log.LOG_FATAL, "Players and indices don't have the same size !")
	}
	for i := 0; i < len(indices); i++ {
		server.Say(playerToString(indices[i], pls[i].GetPlayerName(), pls[i].GetPlayerLevel()))
	}
}

func onSay(emitter *players.Player, args string) {
	server.Say("This was tested")
}

func onMaps(emitter *players.Player, args string) {
	maps := server.GetMaps()
	res := "Available maps: "
	for _, mp := range maps {
		res += mp
		res += ", "
	}
	server.Say(res)
}

func onKick(emitter *players.Player, args string) {
	kickedid, err := strconv.Atoi(args)
	if err != nil {
		return
	}
	kicked, found := players.GetPlayer(kickedid)
	if !found {
		log.Log(log.LOG_VERBOSE, "Player"+args+"does not exist")
	} else {
		if kicked.GetPlayerLevel() < emitter.GetPlayerLevel() {
			log.Log(log.LOG_VERBOSE, "Kicking "+kicked.GetPlayerName())
			server.Say("Kicking" + kicked.GetPlayerName())
			server.Kick(kickedid)
		} else {
			server.Say("Cannot kick a player with higher level than yours !")
			log.Log(log.LOG_VERBOSE, "Cannot kick a player with higher level than yours !")
		}
	}
}
