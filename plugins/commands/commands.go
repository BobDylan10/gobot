package commands

import (
	"strings"
	"testbot/events"
	"testbot/log"

	"testbot/players"

	"testbot/server"
)

type plugininside struct{}

var Plug plugininside = plugininside{}

var initialised = false
var handlers map[string]commandHandler

type commandHandler struct {
	handler func(*players.Player, string)
	level   int
}

//This function allows a plugin to register a command, provided an handler in the form func(arguments)
//However, we need a depencies system to avoid plugins trying to register plugins before it was actually initialised
func RegisterCommand(command string, fn func(*players.Player, string), minlevel int) bool {
	if initialised {
		handlers[command] = commandHandler{fn, minlevel}
		return true
	}
	return false
}

func DeleteCommand(command string) {
	delete(handlers, command)
}

func (p plugininside) Init() chan<- events.Event {
	log.Log(log.LOG_INFO, "Starting plugin Command")

	handlers = make(map[string]commandHandler)
	in := make(chan events.Event)
	go runner(in)
	initialised = true

	RegisterCommand("help", onHelp, 1)
	return in
}

var deps []events.EventType = []events.EventType{events.EVT_CLIENT_SAY}

func (p plugininside) IsDep(e events.EventType) bool {
	for _, v := range deps {
		if v == e {
			return true
		}
	}
	return false
}

func (p plugininside) GetName() string {
	return "commands"
}

func runner(evts <-chan events.Event) {
	//Beware of the deadlock with back !!!
	for {
		evt := <-evts
		switch t := evt.(type) {
		case events.EventSay:
			log.Log(log.LOG_VERBOSE, "Command plugin received", t)
			tmp := strings.SplitN(t.Text, " ", 2) //TODO: put this stuff in a function
			cmd := tmp[0]
			if len(cmd) == 0 {
				continue
			} else {
				if cmd[0] == '!' {
					//This is a command
					cmd = strings.TrimPrefix(cmd, "!")
				} else {
					continue
				}
			}
			args := ""
			if len(tmp) == 2 {
				args = tmp[1]
			}
			if pl, ok := players.GetPlayer(t.Client); ok {
				if handler, ok := handlers[cmd]; ok {
					if handler.level <= pl.GetPlayerLevel() {
						log.Log(log.LOG_VERBOSE, "Executing command", cmd, "with args", args)
						handler.handler(pl, args) //We execute the associated handler
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

func onHelp(emitter *players.Player, args string) {
	allowedCommands := "^0"
	lvl := emitter.GetPlayerLevel()
	for cmd, handlers := range handlers {
		if handlers.level <= lvl {
			allowedCommands += "!" + cmd + ", "
		}
	}
	allowedCommands = strings.TrimSuffix(allowedCommands, ", ")
	server.Say(allowedCommands)
}
