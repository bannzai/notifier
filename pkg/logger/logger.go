package logger

import (
	"fmt"
	"io"
	"os"
	"runtime/debug"
)

type Logger struct {
	io.Writer
}

var sharedLogger = Logger{
	os.Stdout,
}

func Log(v ...interface{}) {
	fmt.Fprint(sharedLogger, v...)
}

func Logf(format string, v ...interface{}) {
	fmt.Fprintf(sharedLogger, format, v...)
}

func Debug() {
	sharedLogger.Write(debug.Stack())
}
