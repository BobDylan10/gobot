package log

import (
	"io"
	"sync"
	"runtime"
	"strconv"
	"os"
	"fmt"
)

var lock     sync.Mutex // ensures atomic writes; protects the following fields
var writer    io.Writer = os.Stdout  // destination for output

type Loglevel int

const (
	LOG_DEBUG Loglevel = iota
	LOG_VERBOSE
	LOG_INFO
	LOG_ERROR //Recoverable error
	LOG_FATAL //Needs the bot to be closed
)

func Init(w io.Writer) {
	writer = w
}

func level(l Loglevel) string {
	switch l {
	case LOG_DEBUG:
		return "DEBUG: "
	case LOG_VERBOSE:
		return "VERBOSE: "
	case LOG_INFO:
		return "INFO: "
	case LOG_ERROR:
		return "ERROR: "
	case LOG_FATAL:
		return "FATAL: "
	default:
		return "LOG: "
	}
}


func Log(lvl Loglevel, a ...interface{}) {
	header := ""
	header += level(lvl)

	_, file, line, ok := runtime.Caller(1)
	finfo := "???: "
	if (ok) {
		finfo = file + " " + strconv.Itoa(line) + ": "
	}
	header += finfo
	fmt.Fprintln(writer, header, a)
}