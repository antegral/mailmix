package Backend

import (
	"context"

	ORM "antegral.net/mailmix/src/Database/Sqlc"
	"github.com/emersion/go-imap/backend"
	"github.com/google/uuid"
)

type ImapUser struct {
	Session uuid.UUID
	Data    ORM.Account
	Query   *ORM.Queries
}

func (User ImapUser) Username() string {
	return User.Data.Username
}

func (User ImapUser) ListMailboxes(subscribed bool) ([]backend.Mailbox, error) {
	ctx := context.Background()

	// TODO: Mailbox (Mailbox implementation: /src/Backend/Mailbox.go)
	_, err := User.Query.GetAllMailBoxInfo(ctx, User.Data.Uuid)
	if err != nil {
		return nil, backend.ErrInvalidCredentials
	}

	return nil, nil
}

func (User ImapUser) GetMailbox(name string) (backend.Mailbox, error) {
	return nil, nil
}

func (User ImapUser) CreateMailbox(name string) error {
	ctx := context.Background()

	_, err := User.Query.CreateMailBox(ctx, ORM.CreateMailBoxParams{
		Uuid:      uuid.New(),
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

	Count, err := User.Query.CountUserOwnedMailBox(ctx, ORM.CountUserOwnedMailBoxParams{
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
