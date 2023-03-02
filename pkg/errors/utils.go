package errors

import (
	"fmt"
	"time"
)

const (
	whenFormat string = "Jan 02 2006 15:04:05.000000"
	sep        string = " ->"
)

func now() string {
	return time.Now().Format(whenFormat)
}

func addWrap(e error, t string) string {
	return fmt.Sprintf("%s%s %s", t, sep, e.Error())
}
