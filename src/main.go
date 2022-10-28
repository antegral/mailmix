package main

import (
	"antegral.net/mailmix/src/Log"
)

func main() {
	if err := Log.Init(4); err != nil {
		panic(err)
	}

	Log.Info.Println("Starting MailMix...")
}
