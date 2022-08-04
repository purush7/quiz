package main

import (
	"strconv"
	"time"

	"github.com/purush7/quiz/functions"
	"github.com/purush7/quiz/inOut"
	log "github.com/sirupsen/logrus"
)

var writer = inOut.Writer

func main() {
	fileName := writer.PutGet("Enter the fileName")
	if fileName == "" {
		fileName = "problems.csv"
	}
	var secs int = 0
	var err error
	duration := new(time.Duration)
	secsString := writer.PutGet("Enter the timer in seconds")
	if secsString != "" {
		secs, err = strconv.Atoi(secsString)
		if err != nil {
			log.Errorln(err)
			duration = nil
		}
	} else {
		duration = nil
	}

	if secs != 0 {
		*duration = time.Duration(secs * int(time.Second))
	}
	functions.StartQuiz(fileName, duration)
}
