package admin

import (
	"testbot/events"
	"testbot/log"

	"testbot/plugins"
	"testbot/plugins/commands"
)

var Plug = plugins.Plugin{Init: Init, Deps: []events.EventType{}}

func Init() chan<- events.Event {
	log.Log(log.LOG_INFO, "Starting plugin Admin")

	if (!commands.RegisterCommand("bonjour", onBonjour, 1)) {
		log.Log(log.LOG_ERROR, "Plugin COMMAND was not initialized !")
	}
	commands.RegisterCommand("iamgod", onIamgod, 1)

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