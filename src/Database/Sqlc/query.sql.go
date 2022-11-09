// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: query.sql

package ORM

import (
	"context"
	"database/sql"
)

const countMailBox = `-- name: CountMailBox :one
SELECT count(*)
FROM MailBox
WHERE name = ?
AND owneruuid = ?
`

type CountMailBoxParams struct {
	Name      string
	Owneruuid string
}

func (q *Queries) CountMailBox(ctx context.Context, arg CountMailBoxParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, countMailBox, arg.Name, arg.Owneruuid)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createAccount = `-- name: CreateAccount :one
INSERT INTO Account (
	uuid,
	username,
	password
) VALUES (
  ?, ?, ?
)
RETURNING uuid, teamuuid, username, password, mailaddress, isquit
`

type CreateAccountParams struct {
	Uuid     string
	Username string
	Password string
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount, arg.Uuid, arg.Username, arg.Password)
	var i Account
	err := row.Scan(
		&i.Uuid,
		&i.Teamuuid,
		&i.Username,
		&i.Password,
		&i.Mailaddress,
		&i.Isquit,
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
RETURNING uuid, region, id, secret, bucketid, bucketprefix
`

type CreateAwsCredentialParams struct {
	Uuid         string
	Region       string
	ID           string
	Secret       string
	Bucketid     string
	Bucketprefix sql.NullString
}

func (q *Queries) CreateAwsCredential(ctx context.Context, arg CreateAwsCredentialParams) (AwsCredential, error) {
	row := q.db.QueryRowContext(ctx, createAwsCredential,
		arg.Uuid,
		arg.Region,
		arg.ID,
		arg.Secret,
		arg.Bucketid,
		arg.Bucketprefix,
	)
	var i AwsCredential
	err := row.Scan(
		&i.Uuid,
		&i.Region,
		&i.ID,
		&i.Secret,
		&i.Bucketid,
		&i.Bucketprefix,
	)
	return i, err
}

const createMail = `-- name: CreateMail :one
INSERT INTO Mail (
	uuid,
	boxuuid
) VALUES (
  ?, ?
)
RETURNING uuid, boxuuid, header, sentat, createdat, size
`

type CreateMailParams struct {
	Uuid    string
	Boxuuid string
}

func (q *Queries) CreateMail(ctx context.Context, arg CreateMailParams) (Mail, error) {
	row := q.db.QueryRowContext(ctx, createMail, arg.Uuid, arg.Boxuuid)
	var i Mail
	err := row.Scan(
		&i.Uuid,
		&i.Boxuuid,
		&i.Header,
		&i.Sentat,
		&i.Createdat,
		&i.Size,
	)
	return i, err
}

const createMailBox = `-- name: CreateMailBox :one
INSERT INTO MailBox (
	uuid,
	name,
	owneruuid
) VALUES (
  ?, ?, ?
)
RETURNING uuid, name, owneruuid
`

type CreateMailBoxParams struct {
	Uuid      string
	Name      string
	Owneruuid string
}

func (q *Queries) CreateMailBox(ctx context.Context, arg CreateMailBoxParams) (MailBox, error) {
	row := q.db.QueryRowContext(ctx, createMailBox, arg.Uuid, arg.Name, arg.Owneruuid)
	var i MailBox
	err := row.Scan(&i.Uuid, &i.Name, &i.Owneruuid)
	return i, err
}

const createSession = `-- name: CreateSession :one
INSERT INTO Session (
	uuid,
	accountuuid
) VALUES (
  ?, ?
)
RETURNING uuid, accountuuid
`

type CreateSessionParams struct {
	Uuid        string
	Accountuuid string
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, createSession, arg.Uuid, arg.Accountuuid)
	var i Session
	err := row.Scan(&i.Uuid, &i.Accountuuid)
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
RETURNING uuid, name, description, credentialuuid
`

type CreateTeamParams struct {
	Uuid           string
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
	)
	return i, err
}

const deleteAllMailInMailBox = `-- name: DeleteAllMailInMailBox :exec
DELETE FROM Mail
WHERE boxuuid = ?
`

func (q *Queries) DeleteAllMailInMailBox(ctx context.Context, boxuuid string) error {
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
	Owneruuid string
}

func (q *Queries) DeleteMailBox(ctx context.Context, arg DeleteMailBoxParams) error {
	_, err := q.db.ExecContext(ctx, deleteMailBox, arg.Name, arg.Owneruuid)
	return err
}

const deleteSessionByUser = `-- name: DeleteSessionByUser :exec
DELETE FROM Session
WHERE accountuuid = ?
`

func (q *Queries) DeleteSessionByUser(ctx context.Context, accountuuid string) error {
	_, err := q.db.ExecContext(ctx, deleteSessionByUser, accountuuid)
	return err
}

const getAccountByUsername = `-- name: GetAccountByUsername :one
SELECT uuid, teamuuid, username, password, mailaddress, isquit FROM Account
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
	)
	return i, err
}

const getAccountByUuid = `-- name: GetAccountByUuid :one
SELECT uuid, teamuuid, username, password, mailaddress, isquit FROM Account
WHERE uuid = ?
`

func (q *Queries) GetAccountByUuid(ctx context.Context, uuid string) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountByUuid, uuid)
	var i Account
	err := row.Scan(
		&i.Uuid,
		&i.Teamuuid,
		&i.Username,
		&i.Password,
		&i.Mailaddress,
		&i.Isquit,
	)
	return i, err
}

const getAllMailBoxInfo = `-- name: GetAllMailBoxInfo :many
SELECT uuid, name, owneruuid FROM MailBox
WHERE owneruuid = ?
`

func (q *Queries) GetAllMailBoxInfo(ctx context.Context, owneruuid string) ([]MailBox, error) {
	rows, err := q.db.QueryContext(ctx, getAllMailBoxInfo, owneruuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MailBox
	for rows.Next() {
		var i MailBox
		if err := rows.Scan(&i.Uuid, &i.Name, &i.Owneruuid); err != nil {
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
SELECT uuid, name, owneruuid FROM MailBox
WHERE name = ?
AND owneruuid = ?
`

type GetOneMailBoxInfoParams struct {
	Name      string
	Owneruuid string
}

func (q *Queries) GetOneMailBoxInfo(ctx context.Context, arg GetOneMailBoxInfoParams) (MailBox, error) {
	row := q.db.QueryRowContext(ctx, getOneMailBoxInfo, arg.Name, arg.Owneruuid)
	var i MailBox
	err := row.Scan(&i.Uuid, &i.Name, &i.Owneruuid)
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
	Owneruuid string
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
