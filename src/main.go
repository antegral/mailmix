package main

import (
	"os"

	DB "antegral.net/mailmix/src/Database"
	"antegral.net/mailmix/src/EnvCheck"
	"antegral.net/mailmix/src/Log"
)

func main() {
	if err := Log.Init(4); err != nil {
		panic(err)
	}

	Log.Info.Println("MailMix IMAP Server")
	Log.Info.Println("Version: at-mm:0.1:221110112439")
	Log.Info.Println("Checking Environment Values...")

	FilePath, err := EnvCheck.GetEnvFilePath()
	Log.Info.Print("Searching: ", FilePath)

	if err != nil {
		panic(err)
	}

	if !EnvCheck.IsFileExists(FilePath) {
		Log.Info.Print("Not found ENV File. Creating File...")
		_, err = os.Create(FilePath)
		if err != nil {
			panic(err)
		}
		EnvCheck.Run()
		Log.Info.Print("ENV File Created. Please set up Environment Values. File Path: ", FilePath)
		os.Exit(1)
	}

	Log.Info.Println("Initializing Database...")

	err = DB.Init()
	if err != nil {
		panic(err)
	}

	Log.Info.Println("Loading Database...")
	_, err = DB.GetDatabase()
	if err != nil {
		panic(err)
	}

	Log.Info.Println("Database Initialized.")
}
