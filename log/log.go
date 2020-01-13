package log

import (
	"io"
	"sync"
	"runtime"
	"strconv"
	"os"
)

var lock     sync.Mutex // ensures atomic writes; protects the following fields
var writer    io.Writer = os.Stdout  // destination for output

type Loglevel int

const (
	LOG_DEBUG Loglevel = iota
	LOG_VERBOSE
	LOG_INFO
	LOG_ERROR
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
	default:
		return "LOG: "
	}
}


func Log(lvl Loglevel, str string) {
	output := make([]byte, 0)
	output = append(output, level(lvl)...)

	_, file, line, ok := runtime.Caller(1)
	finfo := "???: "
	if (ok) {
		finfo = file + " " + strconv.Itoa(line) + ": "
	}
	output = append(output, finfo...)
	output = append(output, str...)
	output = append(output, '\n')
	writer.Write(output) //I hope this is thread-safe (?)
}