package server
import (
	"fmt"
	"testing"
)

// func TestClearLine(t *testing.T) {
// 	Init()
// }

func TestWrap(t *testing.T) {
	l := "This is a veeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeryyyyyyyyyyyyyyyyyyyy looooooooooooooong line"
	r := wrap(l)
	fmt.Println(r)
	fmt.Println(len(r[0]))
}