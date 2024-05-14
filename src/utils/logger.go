package utils

import (
	"log"
	"sync"

	"github.com/fatih/color"
)

type Logger struct{}

var once sync.Once

func GetLoggerInstance() (instance *Logger) {
	once.Do(func() {
		instance = &Logger{}
	})
	return
}

func (l *Logger) Info(msg string) {
	blue := color.New(color.FgBlue).SprintFunc()
	log.Println(blue("[INFO]"), msg)
}

func (l *Logger) Warning(msg string) {
	yellow := color.New(color.FgYellow).SprintFunc()
	log.Println(yellow("[WARNING]"), msg)
}

func (l *Logger) Error(msg string) {
	red := color.New(color.FgRed).SprintFunc()
	log.Println(red("[ERROR]"), msg)
}
