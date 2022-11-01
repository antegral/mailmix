package Server

import (
	"log"

	"antegral.net/mailmix/src/Account"
	"antegral.net/mailmix/src/Log"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/backend"
	"github.com/emersion/go-imap/server"
)

func Listen(Port string, AllowInsecureAuth bool) {
	// Create a new server
	s := server.New(CustomBackend{})
	s.Addr = Port
	// Since we will use this server for testing only, we can allow plain text
	// authentication over unencrypted connections
	s.AllowInsecureAuth = AllowInsecureAuth

	Log.Info.Print("MailMix Server listening on localhost", Port)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

type CustomBackend struct{}

func (be CustomBackend) Login(connInfo *imap.ConnInfo, username, password string) (backend.User, error) {
	return Account.LoginStrategy(connInfo, username, password)
}
