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

func (u *ImapUser) Username() string {
	return u.Data.Username
}

func (u *ImapUser) ListMailboxes(subscribed bool) ([]backend.Mailbox, error) {
	ctx := context.Background()

	// TODO: Mailbox (Mailbox implementation: /src/Backend/Mailbox.go)
	_, err := u.Query.GetAllMailBoxInfo(ctx, u.Data.Uuid)
	if err != nil {
		return nil, backend.ErrInvalidCredentials
	}

	return nil, nil
}

func (u *ImapUser) GetMailbox(name string) (backend.Mailbox, error) {
	return nil, nil
}

func (u *ImapUser) CreateMailbox(name string) error {
	ctx := context.Background()

	_, err := u.Query.CreateMailBox(ctx, ORM.CreateMailBoxParams{
		Uuid:      uuid.New(),
		Name:      name,
		Owneruuid: u.Data.Uuid,
	})
	if err != nil {
		return backend.ErrInvalidCredentials
	}

	return nil
}

func (u *ImapUser) DeleteMailbox(name string) error {
	ctx := context.Background()

	MailBox, err := u.Query.GetOneMailBoxInfo(ctx, ORM.GetOneMailBoxInfoParams{
		Name:      name,
		Owneruuid: u.Data.Uuid,
	})
	if err != nil {
		return backend.ErrNoSuchMailbox
	}

	err = u.Query.DeleteAllMailInMailBox(ctx, MailBox.Uuid)
	if err != nil {
		return backend.ErrNoSuchMailbox
	}

	err = u.Query.DeleteMailBox(ctx, ORM.DeleteMailBoxParams{
		Name:      name,
		Owneruuid: u.Data.Uuid,
	})
	if err != nil {
		return backend.ErrNoSuchMailbox
	}

	return nil
}

func (u *ImapUser) RenameMailbox(existingName, newName string) error {
	ctx := context.Background()

	Count, err := u.Query.CountUserOwnedMailBox(ctx, ORM.CountUserOwnedMailBoxParams{
		Name:      newName,
		Owneruuid: u.Data.Uuid,
	})
	if err != nil {
		return backend.ErrNoSuchMailbox
	}

	if Count != 0 {
		return backend.ErrMailboxAlreadyExists
	}

	err = u.Query.RenameMailBox(ctx, ORM.RenameMailBoxParams{
		Name:      newName,
		Name_2:    existingName,
		Owneruuid: u.Data.Uuid,
	})
	if err != nil {
		return backend.ErrNoSuchMailbox
	}

	return nil
}

func (u *ImapUser) Logout() error {
	ctx := context.Background()

	err := u.Query.DeleteSessionByUser(ctx, u.Data.Uuid)
	if err != nil {
		return backend.ErrInvalidCredentials
	}

	return nil
}
