package EnvCheck

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var EnvKeyList = []string{"MAILMIX_DATABASE_CONNSTRING", "MAILMIX_PASSWORD_SALT"}

func Run() error {
	FilePath, err := GetEnvFilePath()
	if err != nil {
		return err
	}

	if !IsFileExists(FilePath) {
		_, err = os.Create(FilePath)
		if err != nil {
			return err
		}
	}

	File, err := os.OpenFile(FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		return err
	}

	FileInfo, err := os.Stat(FilePath)
	if err != nil {
		return err
	}

	if FileSize := FileInfo.Size(); FileSize <= 1 {
		EnvString := MakeTempleteEnv(EnvKeyList)
		_, err = io.WriteString(File, EnvString)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetEnvFilePath() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Join(pwd, "/", ".env"), nil
}

func IsFileExists(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		fmt.Println("EnvCheck IsFileExists Error!")
		panic(err)
	}
}

func MakeTempleteEnv(Key []string) string {
	StringBuffer := make([]byte, 0)
	for i := 0; i < len(Key); i++ {
		StringBuffer = append(StringBuffer, []byte(fmt.Sprint(Key[i], "=\n"))...)
	}
	return string(StringBuffer[:])
}
