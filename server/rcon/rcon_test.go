package rcon

import (
	"testing"
	"fmt"
)

func TestClearLine(t *testing.T) {
	done := make(chan bool)
	cmds := make(chan string)
	aws := make(chan string)
	go RconRunner(done, cmds, aws)

	cmds<-"kick 0"
	fmt.Println(<-aws)
	done<-true
}