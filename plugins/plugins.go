package plugins

import (
	"testbot/events"

	"testbot/plugins/admin"
	"testbot/plugins/commands"
	"testbot/plugins/stats"
)

var Plugins = map[PluginID]Plugin{
	PLUGIN_CMD: commands.Plug,
	PLUGIN_ADMIN: admin.Plug,
	PLUGIN_STATS: stats.Plug}

type PluginID int

const (
	PLUGIN_CMD PluginID = iota
	PLUGIN_ADMIN
	PLUGIN_STATS
)

func GetPluginOrder() []PluginID {
	return []PluginID{PLUGIN_CMD, PLUGIN_ADMIN, PLUGIN_STATS}
}

type Plugin interface {
	Init() chan<- events.Event
	IsDep(e events.EventType) bool
	GetName() string
}