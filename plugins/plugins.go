package plugins

import (
	"testbot/events"
)

type Plugin int

const (
	PLUGIN_DONE Plugin = iota //Used for saying done with no other message to send
	PLUGIN_BROADCAST Plugin = iota //Used for broadcasting a message to all plugins
	PLUGIN_CMD //Real list starts here
)

type PassEvent struct {
	Dest Plugin
	Evt events.Event
}

