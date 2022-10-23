package Log

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var (
	Verbose *log.Logger
	Warn    *log.Logger
	Info    *log.Logger
	Error   *log.Logger
)

func Init() error {
	Verbose = log.New(os.Stdout, "[VERBOSE] ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(os.Stdout, "[WARNING] ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)

	date := time.Now().Format("2006-01-02 15:04:05")
	pwd, err := os.Getwd()
	if err != nil {
		return err // os.Getwd() 예외처리
	}

	// 경로 포맷팅
	FolderPath := fmt.Sprint(pwd, "/logs")
	FilePath := fmt.Sprint(pwd, "/logs/", date, ".log")

	os.MkdirAll(FolderPath, os.ModePerm)

	if !IsFileExists(FilePath) {
		_, err = os.Create(FilePath)
		if err != nil {
			return err // os.Create() 예외처리
		}
	}

	LogFile, err := os.OpenFile(FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		return err // os.OpenFile() 예외처리
	}

	Writer := io.MultiWriter(LogFile, os.Stdout)

	Verbose.SetOutput(Writer)
	Info.SetOutput(Writer)
	Warn.SetOutput(Writer)
	Error.SetOutput(Writer)

	return nil
}

func IsFileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
