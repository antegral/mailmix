package Log

import (
	"os"
	"testing"
)

func TestLogInit(t *testing.T) {
	Init(os.Stdout, os.Stdout, os.Stdout, os.Stderr)
	Verbose.Println("Test successful.")
	Info.Println("Test successful.")
	Warn.Println("Test successful.")
	Error.Println("Test successful.")
}

func TestUseLogFile(t *testing.T) {
	if err := InitLogFile(); err != nil {
		Error.Println(err)
	}

	Verbose.Println("Test successful.")
	Info.Println("Test successful.")
	Warn.Println("Test successful.")
	Error.Println("Test successful")
}
