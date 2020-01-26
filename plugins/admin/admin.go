package admin

import (
	"testbot/events"
	"testbot/log"

	"testbot/players"

	"testbot/plugins/commands"
)

type plugininside struct{}

var Plug plugininside = plugininside{}

func (p plugininside) Init() chan<- events.Event {
	log.Log(log.LOG_INFO, "Starting plugin Admin")

	if (!commands.RegisterCommand("bonjour", onBonjour, 1)) {
		log.Log(log.LOG_ERROR, "Plugin COMMAND was not initialized !")
	}
	pls := players.GetPlayersOfLevel(100)
	if (len(pls) == 0) {
		commands.RegisterCommand("iamgod", onIamgod, 1) //Only register this command if no superadmin was created
	}
	commands.RegisterCommand("say", onSay, 20)
	commands.RegisterCommand("status", onStatus, 20)
	commands.RegisterCommand("maps", onMaps, 20)
	in := make(chan events.Event)
	go runner(in)
	return in
}

var deps []events.EventType = []events.EventType{}

func (p plugininside) IsDep(e events.EventType) bool{
	for _, v := range deps {
		if (v == e) {
			return true
		}
	}
	return false
}

func (p plugininside) GetName() string {
	return "admin"
}

func runner(evts <-chan events.Event) {
	//Beware of the deadlock with back !!!
	
	log.Log(log.LOG_INFO, "Starting command plugin,", evts)
	for {
		evt := <-evts
		switch t := evt.(type) {
		default:
			log.Log(log.LOG_INFO, "Very weird that we are here, type is ", t)
		}
	}
}