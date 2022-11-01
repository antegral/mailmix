// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package ORM

import (
	"database/sql"
)

type Account struct {
	Uuid     string
	Teamuuid sql.NullString
	Username string
	Password string
}

type AwsCredential struct {
	Uuid         string
	Region       string
	ID           string
	Secret       string
	Bucketid     string
	Bucketprefix sql.NullString
}

type Mail struct {
	Uuid      string
	Owneruuid string
}

type Team struct {
	Uuid           string
	Name           string
	Description    sql.NullString
	Credentialuuid string
}