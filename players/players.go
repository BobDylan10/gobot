package players

import (
	"testbot/events"
	"fmt"
)

type Player struct {
	gid int //The global game ID (unique)
	name string //Current name, used for informative purpose only
	level int //The level of the player on the server
	//data map[string]string //Additional data
}

var players map[int]Player //Players indexed by their current player ID

func CollectEvents(e events.Event) {
	switch t := e.(type) {
	case events.EventClientInfo:
		//Here we must lookup if we already know the user. If yes we grab his info, otherwise we create his entry in the database.
	default:
		fmt.Println("Unexected type ", t)
	}
}


