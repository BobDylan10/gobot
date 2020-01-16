package config

import (
	"testing"
	"fmt"
)

func TestDatabase(t *testing.T) {
	res := readFile("config.json")
	fmt.Println(res["main"])
}