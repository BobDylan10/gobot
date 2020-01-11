package mappings

//This is NOT a PACKAGE, this is used to create mappings for the runner for handy use

import (
	"testbot/events"
	"testbot/plugins"

	"testbot/plugins/commands"
)

//TODO: Create a map with all the runner functions so that main can call them easily
var Runners = map[plugins.Plugin](func (evts <-chan events.Event, back chan<- plugins.PassEvent)){
	plugins.PLUGIN_CMD: commands.Runner}

var deps = map[plugins.Plugin][]events.EventType {
	plugins.PLUGIN_CMD: {events.EVT_CLIENT_SAY} }

func IsDep(p plugins.Plugin, e events.EventType) bool{
	for _, v := range deps[p] {
		if (v == e) {
			return true
		}
	}
	return false
}