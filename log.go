package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Level log level
type Level int

const (
	None Level = iota
	Info
	Error
	Fatal
)

var level2string = [...]string{"", "Info", "Error", "Fatal"}

func (l Level) String() string {
	return level2string[l]
}

func (l Level) Log(format string, v ...interface{}) {
	Log(l, format, v...)
}

func Log(level Level, format string, v ...interface{}) {
	builder := strings.Builder{}
	builder.WriteString(time.Now().UTC().Format("2006-01-02T15:04:05.000Z"))
	builder.WriteByte(' ')
	if level != None {
		builder.WriteString("[" + level.String() + "] ")
	}
	builder.WriteString(format)
	fmt.Fprintf(os.Stderr, builder.String(), v...)
	if level == Fatal {
		os.Exit(1)
	}
}
