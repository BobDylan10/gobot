package commands

import (
	"testing"
	"fmt"
)

//Without initialisation ?
func TestRegisterHandler(t *testing.T) {
	
	if (!RegisterCommand("bs", func(args string) {
		fmt.Println("Oups")
	}, 20)) {
		fmt.Println("Plugin was not initialised : OK")
	}
}