package util

import (
	"encoding/csv"
	"io"

	log "github.com/sirupsen/logrus"
)

// Todo learn from bufio.Scanner and change this
func ReadCSVLine(r io.Reader) chan []string {
	csvReader := csv.NewReader(r)

	ch := make(chan []string)

	go func() {
		for {
			record, err := csvReader.Read()
			if err != nil || err == io.EOF {
				close(ch)
				if err != io.EOF {
					log.Error(err)
				}
				break
			}
			ch <- record
		}
	}()

	return ch

}
