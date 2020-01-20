package plugins

import (
	"testbot/events"

	"testbot/plugins/admin"
	"testbot/plugins/commands"
)

var Plugins = map[PluginID]Plugin{
	PLUGIN_CMD: commands.Plug,
	PLUGIN_ADMIN: admin.Plug}

type PluginID int

const (
	PLUGIN_CMD PluginID = iota
	PLUGIN_ADMIN
)

type Plugin interface {
	Init() chan<- events.Event
	IsDep(e events.EventType) bool
	GetName() string
}