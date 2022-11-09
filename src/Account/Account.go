package Account

import (
	"context"
	"crypto/subtle"

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

	Hash := GetAccountHash(password)
	db := ORM.New(database)

	UserData, err := db.GetAccountByUsername(ctx, username)
	if err != nil {
		return nil, backend.ErrInvalidCredentials
	}

	if res := subtle.ConstantTimeCompare(Hash, []byte(UserData.Password)); res != 0 {
		return nil, backend.ErrInvalidCredentials
	}

	SessionUuid := uuid.New()
	_, err = db.CreateSession(ctx, ORM.CreateSessionParams{
		Uuid:        SessionUuid.String(),
		Accountuuid: UserData.Uuid,
	})

	if err != nil {
		return nil, backend.ErrInvalidCredentials
	}

	return ImapUser{
		Session: SessionUuid,
		Data:    UserData,
		Query:   db,
	}, nil
}

type ImapUser struct {
	Session uuid.UUID
	Data    ORM.Account
	Query   *ORM.Queries
}

func (User ImapUser) Username() string {
	return User.Data.Username
}

func (User ImapUser) ListMailboxes(subscribed bool) ([]backend.Mailbox, error) {
	// ctx := context.Background()

	// // TODO: Mailbox
	// _, err := User.Query.GetAllMailBoxInfo(ctx, User.Data.Uuid)
	// if err != nil {
	// 	return nil, backend.ErrInvalidCredentials
	// }

	return nil, nil
}

func (User ImapUser) GetMailbox(name string) (backend.Mailbox, error) {
	return nil, nil
}

func (User ImapUser) CreateMailbox(name string) error {
	ctx := context.Background()

	_, err := User.Query.CreateMailBox(ctx, ORM.CreateMailBoxParams{
		Uuid:      uuid.New().String(),
		Name:      name,
		Owneruuid: User.Data.Uuid,
	})
	if err != nil {
		return backend.ErrInvalidCredentials
	}

	return nil
}

func (User ImapUser) DeleteMailbox(name string) error {
	ctx := context.Background()

	MailBox, err := User.Query.GetOneMailBoxInfo(ctx, ORM.GetOneMailBoxInfoParams{
		Name:      name,
		Owneruuid: User.Data.Uuid,
	})
	if err != nil {
		return backend.ErrNoSuchMailbox
	}

	err = User.Query.DeleteAllMailInMailBox(ctx, MailBox.Uuid)
	if err != nil {
		return backend.ErrNoSuchMailbox
	}

	err = User.Query.DeleteMailBox(ctx, ORM.DeleteMailBoxParams{
		Name:      name,
		Owneruuid: User.Data.Uuid,
	})
	if err != nil {
		return backend.ErrNoSuchMailbox
	}

	return nil
}

func (User ImapUser) RenameMailbox(existingName, newName string) error {
	ctx := context.Background()

	Count, err := User.Query.CountMailBox(ctx, ORM.CountMailBoxParams{
		Name:      newName,
		Owneruuid: User.Data.Uuid,
	})
	if err != nil {
		return backend.ErrNoSuchMailbox
	}

	if Count != 0 {
		return backend.ErrMailboxAlreadyExists
	}

	err = User.Query.RenameMailBox(ctx, ORM.RenameMailBoxParams{
		Name:      newName,
		Name_2:    existingName,
		Owneruuid: User.Data.Uuid,
	})
	if err != nil {
		return backend.ErrNoSuchMailbox
	}

	return nil
}

func (User ImapUser) Logout() error {
	ctx := context.Background()

	err := User.Query.DeleteSessionByUser(ctx, User.Data.Uuid)
	if err != nil {
		return backend.ErrInvalidCredentials
	}

	return nil
}
