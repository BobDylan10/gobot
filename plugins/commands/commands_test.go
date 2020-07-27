package commands

import (
	"fmt"
	"testbot/players"
	"testing"
)

//Without initialisation ?
func TestRegisterHandler(t *testing.T) {

	if !RegisterCommand("bs", func(emitter *players.Player, args string) {
		fmt.Println("Oups")
	}, 20) {
		fmt.Println("Plugin was not initialised : OK")
	}
}
