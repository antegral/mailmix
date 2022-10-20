package main

import (
	"antegral.net/mailmix/src/Log"
)

func main() {
	if err := Log.InitLogFile(); err != nil {
		Log.Error.Println(err)
	}

	Log.Info.Println("Starting MailMix...")
}
