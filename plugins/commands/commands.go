package commands

import (
	"testbot/events"
	"testbot/plugins"
	"strings"
	"testbot/log"

	"testbot/players"
)

var initialised = false
var handlers map[string] commandHandler

type commandHandler struct {
	handler func(string)
	level int
}

//This function allows a plugin to register a command, provided an handler in the form func(arguments)
//However, we need a depencies system to avoid plugins trying to register plugins before it was actually initialised
func RegisterCommand(command string, fn func(string), minlevel int) bool{
	if (initialised) {
		handlers[command] = commandHandler{fn, minlevel}
		return true
	}
	return false
}

func Runner(evts <-chan events.Event, back chan<- plugins.PassEvent) {
	//Beware of the deadlock with back !!!

	handlers = make(map[string]commandHandler)
	initialised = true
	log.Log(log.LOG_INFO, "Starting command plugin,", evts, back)
	for {
		evt := <-evts
		switch t := evt.(type) {
		case events.EventSay:
			log.Log(log.LOG_DEBUG, "Received a SAY event inside command plugin saying", t.Text)
			tmp := strings.SplitN(t.Text, " ", 2) //TODO: put this stuff in a function
			cmd := tmp[0]
			if (len(cmd) == 0) {
				continue
			} else {
				if (cmd[0] == '!') {
					//This is a command
					cmd = strings.TrimPrefix(cmd, "!")
				} else {
					continue
				}
			}
			args := ""
			if (len(tmp) == 2) {
				args = tmp[1]
			}
			if pl, ok := players.GetPlayer(t.Client); ok {
				if handler, ok := handlers[cmd]; ok {
					if (handler.level <= pl.GetPlayerLevel()) {
						handler.handler(args) //We excute the associated handler
					} else {
						log.Log(log.LOG_VERBOSE, "Player tried to execute command", cmd, "without right priviledges")
						//TODO: Print something
					}
				}
			}
			// fmt.Println("Sending back a command event")
			// back<-plugins.PassEvent{Dest:plugins.PLUGIN_BROADCAST, Evt:events.EventCommand{Client: t.Client, Command: cmd, Args: args}}
		default:
			log.Log(log.LOG_INFO, "Very weird that we are here, type is ", t)
		}
	}
}