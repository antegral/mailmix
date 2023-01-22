// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: query.sql

package ORM

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const countUserOwnedMailBox = `-- name: CountUserOwnedMailBox :one
SELECT count(*)
FROM MailBox
WHERE name = ?
AND owneruuid = ?
`

type CountUserOwnedMailBoxParams struct {
	Name      string
	Owneruuid uuid.UUID
}

func (q *Queries) CountUserOwnedMailBox(ctx context.Context, arg CountUserOwnedMailBoxParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, countUserOwnedMailBox, arg.Name, arg.Owneruuid)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createAccount = `-- name: CreateAccount :one
INSERT INTO Account (
	uuid,
	teamuuid,
	username,
	password,
	mailaddress,
	isquit
) VALUES (
	?, ?, ?, ?, ?, ?
)
RETURNING uuid, teamuuid, username, password, mailaddress, isquit, createdat
`

type CreateAccountParams struct {
	Uuid        uuid.UUID
	Teamuuid    uuid.NullUUID
	Username    string
	Password    string
	Mailaddress string
	Isquit      bool
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount,
		arg.Uuid,
		arg.Teamuuid,
		arg.Username,
		arg.Password,
		arg.Mailaddress,
		arg.Isquit,
	)
	var i Account
	err := row.Scan(
		&i.Uuid,
		&i.Teamuuid,
		&i.Username,
		&i.Password,
		&i.Mailaddress,
		&i.Isquit,
		&i.Createdat,
	)
	return i, err
}

const createAttachment = `-- name: CreateAttachment :one
INSERT INTO Attachment (
	uuid,
	mailuuid,
	cid,
	contenttype
) VALUES (
	?, ?, ?, ?
)
RETURNING uuid, mailuuid, cid, contenttype, createdat
`

type CreateAttachmentParams struct {
	Uuid        uuid.UUID
	Mailuuid    uuid.UUID
	Cid         string
	Contenttype string
}

func (q *Queries) CreateAttachment(ctx context.Context, arg CreateAttachmentParams) (Attachment, error) {
	row := q.db.QueryRowContext(ctx, createAttachment,
		arg.Uuid,
		arg.Mailuuid,
		arg.Cid,
		arg.Contenttype,
	)
	var i Attachment
	err := row.Scan(
		&i.Uuid,
		&i.Mailuuid,
		&i.Cid,
		&i.Contenttype,
		&i.Createdat,
	)
	return i, err
}

const createAwsCredential = `-- name: CreateAwsCredential :one
INSERT INTO AwsCredential (
	uuid,
	region,
	id,
	secret,
	bucketid,
	bucketprefix
) VALUES (
  ?, ?, ?, ?, ?, ?
)
RETURNING uuid, region, id, secret, bucketid, bucketprefix, createdat
`

type CreateAwsCredentialParams struct {
	Uuid         uuid.UUID
	Region       string
	ID           string
	Secret       string
	Bucketid     string
	Bucketprefix sql.NullString
}

func (q *Queries) CreateAwsCredential(ctx context.Context, arg CreateAwsCredentialParams) (Awscredential, error) {
	row := q.db.QueryRowContext(ctx, createAwsCredential,
		arg.Uuid,
		arg.Region,
		arg.ID,
		arg.Secret,
		arg.Bucketid,
		arg.Bucketprefix,
	)
	var i Awscredential
	err := row.Scan(
		&i.Uuid,
		&i.Region,
		&i.ID,
		&i.Secret,
		&i.Bucketid,
		&i.Bucketprefix,
		&i.Createdat,
	)
	return i, err
}

const createEmbeddedFile = `-- name: CreateEmbeddedFile :one
INSERT INTO EmbeddedFile (
	uuid,
	mailuuid,
	contenttype
) VALUES (
	?, ?, ?
)
RETURNING uuid, mailuuid, contenttype, createdat
`

type CreateEmbeddedFileParams struct {
	Uuid        uuid.UUID
	Mailuuid    uuid.UUID
	Contenttype string
}

func (q *Queries) CreateEmbeddedFile(ctx context.Context, arg CreateEmbeddedFileParams) (Embeddedfile, error) {
	row := q.db.QueryRowContext(ctx, createEmbeddedFile, arg.Uuid, arg.Mailuuid, arg.Contenttype)
	var i Embeddedfile
	err := row.Scan(
		&i.Uuid,
		&i.Mailuuid,
		&i.Contenttype,
		&i.Createdat,
	)
	return i, err
}

const createMail = `-- name: CreateMail :one
INSERT INTO Mail (
	uuid,
	boxuuid,
	header,
	sentfrom,
	sentto,
	sentat,
	hash,
	flags, 
	size
) VALUES (
	?, ?, ?, ?, ?, ?, ?, ?, ?
)
RETURNING uid, uuid, boxuuid, header, sentfrom, sentto, sentat, hash, flags, size, createdat
`

type CreateMailParams struct {
	Uuid     uuid.UUID
	Boxuuid  uuid.UUID
	Header   string
	Sentfrom string
	Sentto   string
	Sentat   time.Time
	Hash     string
	Flags    []string
	Size     int32
}

func (q *Queries) CreateMail(ctx context.Context, arg CreateMailParams) (Mail, error) {
	row := q.db.QueryRowContext(ctx, createMail,
		arg.Uuid,
		arg.Boxuuid,
		arg.Header,
		arg.Sentfrom,
		arg.Sentto,
		arg.Sentat,
		arg.Hash,
		pq.Array(arg.Flags),
		arg.Size,
	)
	var i Mail
	err := row.Scan(
		&i.Uid,
		&i.Uuid,
		&i.Boxuuid,
		&i.Header,
		&i.Sentfrom,
		&i.Sentto,
		&i.Sentat,
		&i.Hash,
		pq.Array(&i.Flags),
		&i.Size,
		&i.Createdat,
	)
	return i, err
}

const createMailBox = `-- name: CreateMailBox :one
INSERT INTO MailBox (
	uuid,
	name,
	owneruuid,
  attributes
) VALUES (
  ?, ?, ?, ?
)
RETURNING uuid, name, owneruuid, attributes, messages, recent, unseen, uidnext, uidvalidity, appendlimit, createdat
`

type CreateMailBoxParams struct {
	Uuid       uuid.UUID
	Name       string
	Owneruuid  uuid.UUID
	Attributes []string
}

func (q *Queries) CreateMailBox(ctx context.Context, arg CreateMailBoxParams) (Mailbox, error) {
	row := q.db.QueryRowContext(ctx, createMailBox,
		arg.Uuid,
		arg.Name,
		arg.Owneruuid,
		pq.Array(arg.Attributes),
	)
	var i Mailbox
	err := row.Scan(
		&i.Uuid,
		&i.Name,
		&i.Owneruuid,
		pq.Array(&i.Attributes),
		&i.Messages,
		&i.Recent,
		&i.Unseen,
		&i.Uidnext,
		&i.Uidvalidity,
		&i.Appendlimit,
		&i.Createdat,
	)
	return i, err
}

const createSession = `-- name: CreateSession :one
INSERT INTO Session (
	uuid,
	accountuuid
) VALUES (
	?, ?
)
RETURNING uuid, accountuuid, createdat
`

type CreateSessionParams struct {
	Uuid        uuid.UUID
	Accountuuid uuid.UUID
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, createSession, arg.Uuid, arg.Accountuuid)
	var i Session
	err := row.Scan(&i.Uuid, &i.Accountuuid, &i.Createdat)
	return i, err
}

const createTeam = `-- name: CreateTeam :one
INSERT INTO Team (
	uuid,
	name,
	description,
	credentialuuid
) VALUES (
	?, ?, ?, ?
)
RETURNING uuid, name, description, credentialuuid, createdat
`

type CreateTeamParams struct {
	Uuid           uuid.UUID
	Name           string
	Description    sql.NullString
	Credentialuuid string
}

func (q *Queries) CreateTeam(ctx context.Context, arg CreateTeamParams) (Team, error) {
	row := q.db.QueryRowContext(ctx, createTeam,
		arg.Uuid,
		arg.Name,
		arg.Description,
		arg.Credentialuuid,
	)
	var i Team
	err := row.Scan(
		&i.Uuid,
		&i.Name,
		&i.Description,
		&i.Credentialuuid,
		&i.Createdat,
	)
	return i, err
}

const deleteAllMailInMailBox = `-- name: DeleteAllMailInMailBox :exec
DELETE FROM Mail
WHERE boxuuid = ?
`

func (q *Queries) DeleteAllMailInMailBox(ctx context.Context, boxuuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAllMailInMailBox, boxuuid)
	return err
}

const deleteMailBox = `-- name: DeleteMailBox :exec
DELETE FROM MailBox
WHERE name = ?
AND owneruuid = ?
`

type DeleteMailBoxParams struct {
	Name      string
	Owneruuid uuid.UUID
}

func (q *Queries) DeleteMailBox(ctx context.Context, arg DeleteMailBoxParams) error {
	_, err := q.db.ExecContext(ctx, deleteMailBox, arg.Name, arg.Owneruuid)
	return err
}

const deleteSessionByUser = `-- name: DeleteSessionByUser :exec
DELETE FROM Session
WHERE accountuuid = ?
`

func (q *Queries) DeleteSessionByUser(ctx context.Context, accountuuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteSessionByUser, accountuuid)
	return err
}

const getAccountByUsername = `-- name: GetAccountByUsername :one
SELECT uuid, teamuuid, username, password, mailaddress, isquit, createdat FROM Account
WHERE username = ?
`

func (q *Queries) GetAccountByUsername(ctx context.Context, username string) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountByUsername, username)
	var i Account
	err := row.Scan(
		&i.Uuid,
		&i.Teamuuid,
		&i.Username,
		&i.Password,
		&i.Mailaddress,
		&i.Isquit,
		&i.Createdat,
	)
	return i, err
}

const getAccountByUuid = `-- name: GetAccountByUuid :one
SELECT uuid, teamuuid, username, password, mailaddress, isquit, createdat FROM Account
WHERE uuid = ?
`

func (q *Queries) GetAccountByUuid(ctx context.Context, uuid uuid.UUID) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountByUuid, uuid)
	var i Account
	err := row.Scan(
		&i.Uuid,
		&i.Teamuuid,
		&i.Username,
		&i.Password,
		&i.Mailaddress,
		&i.Isquit,
		&i.Createdat,
	)
	return i, err
}

const getAllMailBoxInfo = `-- name: GetAllMailBoxInfo :many
SELECT uuid, name, owneruuid, attributes, messages, recent, unseen, uidnext, uidvalidity, appendlimit, createdat FROM MailBox
WHERE owneruuid = ?
`

func (q *Queries) GetAllMailBoxInfo(ctx context.Context, owneruuid uuid.UUID) ([]Mailbox, error) {
	rows, err := q.db.QueryContext(ctx, getAllMailBoxInfo, owneruuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Mailbox
	for rows.Next() {
		var i Mailbox
		if err := rows.Scan(
			&i.Uuid,
			&i.Name,
			&i.Owneruuid,
			pq.Array(&i.Attributes),
			&i.Messages,
			&i.Recent,
			&i.Unseen,
			&i.Uidnext,
			&i.Uidvalidity,
			&i.Appendlimit,
			&i.Createdat,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOneMailBoxInfo = `-- name: GetOneMailBoxInfo :one
SELECT uuid, name, owneruuid, attributes, messages, recent, unseen, uidnext, uidvalidity, appendlimit, createdat FROM MailBox
WHERE name = ?
AND owneruuid = ?
`

type GetOneMailBoxInfoParams struct {
	Name      string
	Owneruuid uuid.UUID
}

func (q *Queries) GetOneMailBoxInfo(ctx context.Context, arg GetOneMailBoxInfoParams) (Mailbox, error) {
	row := q.db.QueryRowContext(ctx, getOneMailBoxInfo, arg.Name, arg.Owneruuid)
	var i Mailbox
	err := row.Scan(
		&i.Uuid,
		&i.Name,
		&i.Owneruuid,
		pq.Array(&i.Attributes),
		&i.Messages,
		&i.Recent,
		&i.Unseen,
		&i.Uidnext,
		&i.Uidvalidity,
		&i.Appendlimit,
		&i.Createdat,
	)
	return i, err
}

const renameMailBox = `-- name: RenameMailBox :exec
UPDATE MailBox
SET name = ?
WHERE name = ?
AND owneruuid = ?
`

type RenameMailBoxParams struct {
	Name      string
	Name_2    string
	Owneruuid uuid.UUID
}

func (q *Queries) RenameMailBox(ctx context.Context, arg RenameMailBoxParams) error {
	_, err := q.db.ExecContext(ctx, renameMailBox, arg.Name, arg.Name_2, arg.Owneruuid)
	return err
}

const setAccountPassword = `-- name: SetAccountPassword :exec
UPDATE Account
SET password = ?
WHERE username = ?
`

type SetAccountPasswordParams struct {
	Password string
	Username string
}

func (q *Queries) SetAccountPassword(ctx context.Context, arg SetAccountPasswordParams) error {
	_, err := q.db.ExecContext(ctx, setAccountPassword, arg.Password, arg.Username)
	return err
}

const updateMailBox = `-- name: UpdateMailBox :exec
UPDATE MailBox
SET name = ?
WHERE name = ?
AND owneruuid = ?
`

type UpdateMailBoxParams struct {
	Name      string
	Name_2    string
	Owneruuid uuid.UUID
}

func (q *Queries) UpdateMailBox(ctx context.Context, arg UpdateMailBoxParams) error {
	_, err := q.db.ExecContext(ctx, updateMailBox, arg.Name, arg.Name_2, arg.Owneruuid)
	return err
}
