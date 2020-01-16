package admin

import (
	"testbot/events"
	"testbot/log"

	"testbot/plugins"

	"testbot/plugins/commands"
)

func Runner(evts <-chan events.Event, back chan<- plugins.PassEvent) {
	//Beware of the deadlock with back !!!
	if (!commands.RegisterCommand("bonjour", onBonjour, 1)) {
		log.Log(log.LOG_ERROR, "Plugin COMMAND was not initialized !")
	}
	log.Log(log.LOG_INFO, "Starting command plugin,", evts, back)
	for {
		evt := <-evts
		switch t := evt.(type) {
		default:
			log.Log(log.LOG_INFO, "Very weird that we are here, type is ", t)
		}
	}
}