package iourt43

import (
	"fmt"
)

type parseError struct {
	line string
}

func (e *parseError) Error() string {
    return fmt.Sprintf("Error parsing line : \"%s\"", e.line)
}