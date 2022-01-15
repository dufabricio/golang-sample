package main

import (
	"fmt"
	"time"
)

const (
	levelInfo = "INFO"
	levelWarning = "WARN"
	levelError = "ERROR"
)

type logEntry struct {
	time time.Time
	severity string
	message string
}

var logChannel = make(chan logEntry, 50);

func main() {
	go logger()
	logChannel <- logEntry{time.Now(), levelInfo, "App is starting"}
	time.Sleep(5000 * time.Millisecond)
	logChannel <- logEntry{time.Now(), levelInfo, "App is shutting down"}
	time.Sleep(100 * time.Millisecond)
}

func logger() { 
	for entry := range logChannel {
		 fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02-T15:04:05"), entry.severity, entry.message)
	}
}