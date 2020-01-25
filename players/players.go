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
	guid string //The game unique ID
	toBeDeleted bool
	//Map for all extra attributes ? Seems goooooooood
}

var players map[int]*player //Players indexed by their current player ID

func Init() {
	players = make(map[int]*player)
	initTable()
}

func CollectEvents(e events.Event) {
	//First, we remove the players to be deleted
	for i, pl := range players {
		if pl.toBeDeleted {
			delete(players, i)
		}
	}
	//Then we can do our work normally
	switch t := e.(type) {
	case events.EventClientInfo:
		//We check that the player is not already inside the connected players
		if pl, ok := players[t.Client]; ok {
			//We check that the guid is corresponding
			if (pl.guid != t.Data["cl_guid"]) {
				log.Log(log.LOG_ERROR, "A player id was seen with a GUID different than in the database")
			}
		} else {
			//We had the player to the connected player list
			//Here we must lookup if we already know the user. If yes we grab his info, otherwise we create his entry in the database.
			if guid, present := t.Data["cl_guid"]; present {
				if name, present := t.Data["name"]; present {
					pl := &player{}
					getPlayer(guid, name, pl)
					log.Log(log.LOG_DEBUG, "Player with Database id", pl.did, ", name", pl.name)
					players[t.Client] = pl
					pl.newConnection() //This must be called only once we checked that he is not already in the connected players
				}
			} else {
				//It's a bot ?
			}
		}
	case events.EventClientDisconnect:
		players[t.Client].toBeDeleted = true
	default:
		log.Log(log.LOG_DEBUG, "Unexpected type", t)
	}
}

func GetPlayer(clientid int) (*player, bool) {
	if pl, ok := players[clientid]; ok {
		return pl, true
	}
	return &player{}, false
}

func (pl *player) GetPlayerLevel() int {
	return pl.level
}

