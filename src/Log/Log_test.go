package Log

import (
	"testing"
)

func TestLogInit(t *testing.T) {
	Init()
	Verbose.Println("Test successful.")
	Info.Println("Test successful.")
	Warn.Println("Test successful.")
	Error.Println("Test successful.")
}
