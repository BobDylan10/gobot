package plugins

import (
	"testbot/events"
)

type PluginID int

const (
	PLUGIN_CMD PluginID = iota
	PLUGIN_ADMIN
)

type Plugin struct {
	Init func() chan<- events.Event
	Deps []events.EventType
}