package players

import (
	"testbot/events"
	"fmt"
	"time"
)

type Player struct {
	cgid int //Current game ID
	did int //Database ID
	name string //Current name, used for informative purpose only
	level int //The level of the player on the server
	//data map[string]string //Additional data
	isBot bool
	connections int
	lastConnection time.Time
	//Map for all extra attributes ? Seems goooooooood
}

var players map[int]Player //Players indexed by their current player ID

func Init() {
	players = make(map[int]Player)
}

func CollectEvents(e events.Event) {
	switch t := e.(type) {
	case events.EventClientInfo:
		//Here we must lookup if we already know the user. If yes we grab his info, otherwise we create his entry in the database.
		if guid, present := t.Data["cl_guid"]; present {
			if name, present := t.Data["name"]; present {
			fmt.Println("Found GUID:", guid)
			getPlayerDid(guid, name)
			}
		} else {
			//It's a bot ?
		}
	default:
		fmt.Println("Unexected type ", t)
	}
}


