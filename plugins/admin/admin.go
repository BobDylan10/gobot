package admin

import (
	"testbot/events"
	"testbot/log"

	"testbot/plugins/commands"
)

func Runner(evts <-chan events.Event) {
	//Beware of the deadlock with back !!!
	if (!commands.RegisterCommand("bonjour", onIamgod, 1)) {
		log.Log(log.LOG_ERROR, "Plugin COMMAND was not initialized !")
	}
	commands.RegisterCommand("bonjour", onIamgod, 1)
	log.Log(log.LOG_INFO, "Starting command plugin,", evts)
	for {
		evt := <-evts
		switch t := evt.(type) {
		default:
			log.Log(log.LOG_INFO, "Very weird that we are here, type is ", t)
		}
	}
}