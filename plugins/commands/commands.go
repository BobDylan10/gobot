package commands

import (
	"testbot/events"
	"testbot/plugins"
	"strings"
	"fmt"
)

var initialised = false
var handlers map[string]func(string, string)

//This function allows a plugin to register a command, provided an handler in the form func(command, arguments)
//However, we need a depencies system to avoid plugins trying to register plugins before it was actually initialised
func RegisterCommand(command string, fn func(string, string)) bool{
	if (initialised) {
		handlers[command] = fn
		return true
	}
	return false
}

func Runner(evts <-chan events.Event, back chan<- plugins.PassEvent) {
	//Beware of the deadlock with back !!!

	handlers = make(map[string]func(string, string))
	initialised = true
	fmt.Println("Starting command plugin,", evts, back)
	for {
		evt := <-evts
		switch t := evt.(type) {
		case events.EventSay:
			fmt.Println("Received a SAY event inside command plugin !!!")
			tmp := strings.SplitN(t.Text, " ", 2)
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
			if handler, ok := handlers[cmd]; ok {
				handler(cmd, args) //We excute the associated handler
			}
			// fmt.Println("Sending back a command event")
			// back<-plugins.PassEvent{Dest:plugins.PLUGIN_BROADCAST, Evt:events.EventCommand{Client: t.Client, Command: cmd, Args: args}}
		default:
			fmt.Println("Very weird that we are here, type is ", t)
		}
	}
}