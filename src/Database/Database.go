package Database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

var createTableQuery = `
CREATE TABLE IF NOT EXISTS AwsCredential (
	Uuid					TEXT NOT NULL,
	Region				TEXT NOT NULL,
	Id						TEXT NOT NULL UNIQUE,
	Secret				TEXT NOT NULL,
	BucketId			TEXT NOT NULL,
	BucketPrefix	TEXT,
  CreatedAt			TIMESTEMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(Uuid)
);

CREATE TABLE IF NOT EXISTS Team (
	Uuid						TEXT NOT NULL,
  Name						TEXT NOT NULL UNIQUE,
  Description			TEXT,
	CredentialUuid	TEXT NOT NULL UNIQUE,
  CreatedAt				TIMESTEMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(Uuid),
	FOREIGN KEY(CredentialUuid) REFERENCES AwsCredential(Uuid) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS Account (
	Uuid				TEXT NOT NULL,
	TeamUuid		TEXT,
	Username		TEXT NOT NULL UNIQUE,
	Password		TEXT NOT NULL,
	MailAddress	TEXT NOT NULL UNIQUE,
	IsQuit      INTEGER NOT NULL DEFAULT 0,
  CreatedAt		TIMESTEMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY(Uuid),
	FOREIGN KEY(TeamUuid) REFERENCES Team(Uuid) ON DELETE SET NULL ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS MailBox (
	Uuid				TEXT NOT NULL,
	Name				TEXT NOT NULL,
	OwnerUuid		TEXT NOT NULL,
  CreatedAt		TIMESTEMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(Uuid)
);

CREATE TABLE Mail (
	Uuid				TEXT NOT NULL,
  BoxUuid     TEXT NOT NULL,
	CreatedAt		TIMESTEMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(Uuid)
);

CREATE TABLE IF NOT EXISTS Session (
	Uuid        TEXT NOT NULL,
	AccountUuid TEXT NOT NULL,
  CreatedAt		TIMESTEMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(Uuid)
);
`

func Init() error {
	DBFilePath, err := GetDatabasePath()
	if err != nil {
		return err
	}

	if _, err := os.Create(DBFilePath); err != nil {
		return err
	}

	Database, err := GetDatabase()
	if err != nil {
		return err
	}

	if _, err = Database.Exec(createTableQuery); err != nil {
		return err
	}

	return nil
}

func GetDatabase() (*sql.DB, error) {
	DBFilePath, err := GetDatabasePath()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", DBFilePath)
	if err != nil {
		return nil, err
	}
	return db, err
}

func GetDatabasePath() (string, error) {
	// 현재 경로 가져오기
	pwd, err := os.Getwd()
	if err != nil {
		return "", err // os.Getwd() 예외처리
	}

	// 경로 포맷팅
	return fmt.Sprint(pwd, "/database.db"), nil
}

func GetUser(Uuid uuid.UUID) {
}
