package log


import (
	"testing"
	"os"
)

func TestLog(t *testing.T) {
	Init(os.Stdout)
	Log(LOG_DEBUG, "aaa")
}