package inOut

import (
	"os"

	"github.com/purush7/quiz/constants"
	"github.com/purush7/quiz/util"
)

type inOutInterface interface {
	PutGet(string) string
	Put(string)
}

var Writer inOutInterface

func init() {
	value := os.Getenv("SETUP")

	if value == constants.FILE {
		fileName := os.Getenv("FILENAME")
		if fileName == "" || util.StatFile(fileName) {
			value = constants.TERMINAL // make it default
		}
	}

	switch value {
	case constants.API:
		APITextInit()

	case constants.FILE:

	default:
		Writer = terminalInit()
	}
}
