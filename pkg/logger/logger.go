package logger

import (
	"fmt"
	"io"
	"os"
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
