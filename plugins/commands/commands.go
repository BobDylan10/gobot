package commands

import (
	"testbot/plugins"

	"testbot/events"
	"strings"
	"testbot/log"

	"testbot/players"

	"testbot/server"
)

var Plug = plugins.Plugin{Init: Init, Deps: []events.EventType{events.EVT_CLIENT_SAY}}

var initialised = false
var handlers map[string] commandHandler

type commandHandler struct {
	handler func(int, string)
	level int
}

//This function allows a plugin to register a command, provided an handler in the form func(arguments)
//However, we need a depencies system to avoid plugins trying to register plugins before it was actually initialised
func RegisterCommand(command string, fn func(int, string), minlevel int) bool{
	if (initialised) {
		handlers[command] = commandHandler{fn, minlevel}
		return true
	}
	return false
}

func Init() chan<- events.Event{
	log.Log(log.LOG_INFO, "Starting plugin Command")

	handlers = make(map[string]commandHandler)
	RegisterCommand("help", onHelp, 1)

	in := make(chan events.Event)
	go runner(in)
	initialised = true
	return in
}

func runner(evts <-chan events.Event) {
	//Beware of the deadlock with back !!!
	for {
		evt := <-evts
		switch t := evt.(type) {
		case events.EventSay:
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
						log.Log(log.LOG_VERBOSE, "Executing command", cmd, "with args", args)
						handler.handler(t.Client, args) //We execute the associated handler
					} else {
						log.Log(log.LOG_VERBOSE, "Player tried to execute command", cmd, "without right priviledges")
						//TODO: Print something
					}
				}
			}
		default:
			log.Log(log.LOG_INFO, "Very weird that we are here, type is ", t)
		}
	}
}

func onHelp(id int, args string) {
	pl, ok := players.GetPlayer(id)
	allowedCommands := ""
	if (ok) {
		lvl := pl.GetPlayerLevel()
		for cmd, handlers := range handlers {
			if (handlers.level <= lvl) {
				allowedCommands += cmd + ", "
			}
		}
		server.Say(allowedCommands)
	}
}