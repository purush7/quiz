package functions

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/purush7/quiz/util"
	log "github.com/sirupsen/logrus"
)

//Todo use custom logger
// var log = logger.Init()

func getAns() string {
	var str string
	fmt.Scanf("%s", &str)
	return str
}

func pushResult(str string) {
	fmt.Println(str)
}

func StartQuiz(fileName string, timer *time.Duration) {
	present := util.StatFile(fileName)
	if !present {
		log.Errorf("%s file isn't present", fileName)
		return
	}
	file, err := os.Open(fileName)
	if err != nil {
		log.Errorln(err)
		return
	}
	ansCh := make(chan string)

	//Get input
	go func() {
		for {
			ansCh <- getAns()
		}
	}()

	result := askQuestions(timer, file, os.Stdout, ansCh)

	pushResult(result)
}

func askQuestions(timer *time.Duration, r io.Reader, w io.Writer, ansCh chan string) string {
	score := 0
	ind := 0
	quesCh := util.ReadCSVLine(r)
	for record := range quesCh {
		if len(record) != 2 {
			log.Errorln("error length of record %v isn't 2 as expected", record)
			close(quesCh)
			break
		}
		fmt.Fprintf(w, "what %s, sir?\n", record[0])
		if timer != nil {
			go func() {
				time.Sleep(*timer)
				ansCh <- ""
			}()
		}
		ans := <-ansCh

		if ans == record[1] {
			score++
		}
		ind++
	}
	close(ansCh)
	return fmt.Sprintf("You got %d correct out of %d", score, ind)
}
