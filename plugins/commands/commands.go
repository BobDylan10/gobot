package commands

import (
	"testbot/events"
	"testbot/plugins"
	"strings"
	"fmt"
)

func Runner(evts <-chan events.Event, back chan<- plugins.PassEvent) {
	//Beware of the deadlock with back !!!
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
			fmt.Println("Sending back a command event")
			back<-plugins.PassEvent{Dest:plugins.PLUGIN_BROADCAST, Evt:events.EventCommand{Client: t.Client, Command: cmd, Args: args}}
		default:
			fmt.Println("Very weird that we are here, type is ", t)
		}
	}
}