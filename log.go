package main

import (
	"fmt"
	. "github.com/logrusorgru/aurora/v4"
	"log"
	"os"
)

func logError(message string, v ...interface{}) {
	log.Println(Bold(Red("ERROR:")), fmt.Sprintf(message, v...))
}

func logAndExit(exitCode int, message string, v ...interface{}) {
	log.Println(Bold(Red("FATAL:")), fmt.Sprintf(message, v...))
	os.Exit(exitCode)
}

func logWarning(message string, v ...interface{}) {
	log.Println(Bold(Yellow("WARNING:")), fmt.Sprintf(message, v...))
}
