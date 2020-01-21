package config

import (
	"testing"
	"fmt"
)

func TestDatabase(t *testing.T) {
	loadCfg("config.json")
	s := NewCfg("main")
	fmt.Println(s.GetString("prefix", "default"))
	fmt.Println(s.GetFloat("delay", 123))
}