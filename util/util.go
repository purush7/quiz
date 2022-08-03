package util

import (
	"log"
	"os"
)

func StatFile(fileName string) bool {
	if _, err := os.Stat(fileName); err != nil {
		log.Printf("file %s isn't present\n", fileName)
		return false
	}
	return true
}
