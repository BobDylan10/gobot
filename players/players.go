package players

import (
	"testbot/events"
	"time"
	"testbot/log"
)

type player struct {
	did int //Database ID
	name string //Current name, used for informative purpose only
	level int //The level of the player on the server
	//data map[string]string //Additional data
	isBot bool
	connections int
	lastConnection time.Time
	//Map for all extra attributes ? Seems goooooooood
}

var players map[int]player //Players indexed by their current player ID

func Init() {
	players = make(map[int]player)
}

func CollectEvents(e events.Event) {
	switch t := e.(type) {
	case events.EventClientInfo:
		//TODO: First we check if the player is not already connected !
		//Here we must lookup if we already know the user. If yes we grab his info, otherwise we create his entry in the database.
		if guid, present := t.Data["cl_guid"]; present {
			if name, present := t.Data["name"]; present {
				pl := &player{}
				getPlayer(guid, name, pl)
				log.Log(log.LOG_DEBUG, "Player with Database id", pl.did, ", name", pl.name)
				players[t.Client] = *pl
			}
		} else {
			//It's a bot ?
		}
	default:
		log.Log(log.LOG_INFO, "Unexpected type", t)
	}
}

func GetPlayer(clientid int) (player, bool) {
	if pl, ok := players[clientid]; ok {
		return pl, true
	}
	return player{}, false
}

func (pl *player) GetPlayerLevel() int {
	return pl.level
}

