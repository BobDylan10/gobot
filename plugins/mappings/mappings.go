package mappings

//This is NOT a PACKAGE, this is used to create mappings for the runner for handy use

import (
	"testbot/events"

	"testbot/plugins"

	"testbot/plugins/commands"
	"testbot/plugins/admin"
)

var Plugins = map[plugins.PluginID]plugins.Plugin{
	plugins.PLUGIN_CMD: commands.Plug,
	plugins.PLUGIN_ADMIN: admin.Plug}

func IsDep(p plugins.PluginID, e events.EventType) bool{
	for _, v := range Plugins[p].Deps {
		if (v == e) {
			return true
		}
	}
	return false
}