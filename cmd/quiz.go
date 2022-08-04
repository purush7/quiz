package main

import (
	"fmt"
	"time"

	"github.com/purush7/quiz/functions"
)

func main() {
	fmt.Println("Enter the fileName")
	var fileName string
	fmt.Scanf("%s", &fileName)
	if fileName == "" {
		fileName = "problems.csv"
	}
	var secs int = 0
	timer := new(time.Duration)
	fmt.Println("Enter the timer in seconds")
	fmt.Scanf("%d", &secs)
	if secs == 0 {
		timer = nil
	} else {
		*timer = time.Duration(secs * int(time.Second))
	}
	functions.StartQuiz(fileName, timer)
}
