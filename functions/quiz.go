package functions

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/purush7/quiz/inOut"
	"github.com/purush7/quiz/util"
	log "github.com/sirupsen/logrus"
)

//Todo use custom logger
// var log = logger.Init()

var writer = inOut.Writer

func StartQuiz(fileName string, duration *time.Duration) {
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

	result := askQuestions(duration, file, os.Stdout)

	writer.Put(result)
}

func askQuestions(duration *time.Duration, r io.Reader, w io.Writer) string {
	score := 0
	ind := 0
	quesCh := util.ReadCSVLine(r)

	ansCh := make(chan string)

	for record := range quesCh {
		if len(record) != 2 {
			log.Errorln("error length of record %v isn't 2 as expected", record)
			close(quesCh)
			break
		}

		if duration != nil {
			timer := time.NewTimer(*duration)
			go func(record []string) {
				ansCh <- writer.PutGet(fmt.Sprintf("what %s, sir?", record[0]))
			}(record)

			select {
			case <-timer.C:
				writer.Put("Timeout")
			case ans := <-ansCh:
				if ans == record[1] {
					score++
				}
			}
			ind++
			continue
		}
		ans := writer.PutGet(fmt.Sprintf("what %s, sir?", record[0]))
		if ans == record[1] {
			score++
		}
		ind++
	}
	close(ansCh)
	return fmt.Sprintf("You got %d correct out of %d", score, ind)
}
