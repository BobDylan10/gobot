package admin

import (
	"testbot/events"
	"testbot/log"

	"testbot/players"

	"testbot/plugins"
	"testbot/plugins/commands"
)

var Plug = plugins.Plugin{Init: Init, Deps: []events.EventType{}}

func Init() chan<- events.Event {
	log.Log(log.LOG_INFO, "Starting plugin Admin")

	if (!commands.RegisterCommand("bonjour", onBonjour, 1)) {
		log.Log(log.LOG_ERROR, "Plugin COMMAND was not initialized !")
	}
	pls := players.GetPlayersOfLevel(100)
	if (len(pls) == 0) {
		commands.RegisterCommand("iamgod", onIamgod, 1) //Only register this command if no superadmin was created
	}
	commands.RegisterCommand("say", onSay, 20)
	in := make(chan events.Event)
	go runner(in)
	return in
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