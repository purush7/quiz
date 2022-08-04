package io

import (
	"io"
	"os"

	"github.com/purush7/quiz/constants"
	"github.com/purush7/quiz/util"
)

type writerStruct struct {
	writer io.Writer
}

var writer writerStruct

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

	}
}
