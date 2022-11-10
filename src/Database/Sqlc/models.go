// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package ORM

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Uuid        uuid.UUID
	Teamuuid    uuid.NullUUID
	Username    string
	Password    string
	Mailaddress string
	Isquit      bool
	Createdat   time.Time
}

type Attachment struct {
	Uuid        uuid.UUID
	Mailuuid    uuid.UUID
	Cid         string
	Contenttype string
	Createdat   time.Time
}

type Awscredential struct {
	Uuid         uuid.UUID
	Region       string
	ID           string
	Secret       string
	Bucketid     string
	Bucketprefix sql.NullString
	Createdat    time.Time
}

type Embeddedfile struct {
	Uuid        uuid.UUID
	Mailuuid    uuid.UUID
	Contenttype string
	Createdat   time.Time
}

type Mail struct {
	Uuid      uuid.UUID
	Boxuuid   uuid.UUID
	Header    string
	Sentfrom  string
	Sentto    string
	Sentat    time.Time
	Content   string
	Flags     string
	Size      int32
	Createdat time.Time
}

type Mailbox struct {
	Uuid       uuid.UUID
	Name       string
	Owneruuid  uuid.UUID
	Attributes []string
	Createdat  time.Time
}

type Session struct {
	Uuid        uuid.UUID
	Accountuuid uuid.UUID
	Createdat   time.Time
}

type Team struct {
	Uuid           uuid.UUID
	Name           string
	Description    sql.NullString
	Credentialuuid string
	Createdat      time.Time
}
