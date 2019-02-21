package server

import (
	"fmt"
	"github.com/pkg/errors"
)

type SuperComplexLogger interface {
	Log(message string)
	SuperLog(message string) (string, int, int, error)
}

type MattsLogger struct {
}

func (MattsLogger) Log(message string) {
	fmt.Println(message)
}

func (MattsLogger) SuperLog(message string) (string, int, int, error) {
	switch message {
	case "hello":
		return "hello", 1, 1, nil
	case "this logger is stupid":
		return "I agree", 0, 0, nil
	default:
		return "", 0, 0, errors.New("what a terrible logger")
	}
}
