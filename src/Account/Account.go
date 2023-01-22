package Account

import (
	"context"
	"crypto/subtle"

	Backend "antegral.net/mailmix/src/Backend"
	"antegral.net/mailmix/src/Database"
	ORM "antegral.net/mailmix/src/Database/Sqlc"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/backend"
	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"
)

type AwsCredential struct {
	Region       string
	ApiId        string
	ApiKey       string
	BucketId     string
	BucketPrefix string
}

func GetAccountHash(Plain string) []byte {
	return argon2.IDKey([]byte(Plain), []byte("asdf"), 1, 64*1024, 4, 32)
}

func LoginStrategy(connInfo *imap.ConnInfo, username, password string) (backend.User, error) {
	ctx := context.Background()
	database, err := Database.GetDatabase()
	if err != nil {
		return nil, backend.ErrInvalidCredentials
	}

	// Hashing the entered plaintext password
	Hash := GetAccountHash(password)

	// Import ORM
	Queries := ORM.New(database)

	// Getting user information stored in the database
	UserData, err := Queries.GetAccountByUsername(ctx, username)
	if err != nil {
		// Hash mismatch
		return nil, backend.ErrInvalidCredentials
	}

	// Compare Password Hash
	if res := subtle.ConstantTimeCompare(Hash, []byte(UserData.Password)); res != 0 {
		return nil, backend.ErrInvalidCredentials
	}

	// Create a new session UUID
	SessionUuid := uuid.New()
	_, err = Queries.CreateSession(ctx, ORM.CreateSessionParams{
		Uuid:        SessionUuid,
		Accountuuid: UserData.Uuid,
	})

	if err != nil {
		// Database processing error
		return nil, backend.ErrInvalidCredentials
	}

	return Backend.ImapUser{
		Session: SessionUuid,
		Data:    UserData,
		Query:   Queries,
	}, nil
}
