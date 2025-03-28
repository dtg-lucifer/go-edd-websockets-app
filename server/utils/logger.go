package utils

import (
	"log"
	"os"
)

type Logger struct{}

var iLogger = log.New(os.Stdout, "[Info] - ", log.LstdFlags)
var eLogger = log.New(os.Stderr, "[Error] - ", log.LstdFlags)

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Info(v any) {
	iLogger.Printf("%v", v)
}

func (l *Logger) Err(v any) {
	eLogger.Fatalf("%v", v)
}
