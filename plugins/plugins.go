package plugins

import (
	"testbot/events"
)

type Plugin int

const (
	PLUGIN_DONE Plugin = iota //Used for saying done with no other message to send
	PLUGIN_BROADCAST //Used for broadcasting a message to all plugins - TODO: useless ? yes i think so
	PLUGIN_CMD //Real list starts here
	PLUGIN_ADMIN
)

type PassEvent struct {
	Dest Plugin
	Evt events.Event
}

