CREATE TABLE AwsCredential (
	Uuid					TEXT NOT NULL,
	Region				TEXT NOT NULL,
	Id						TEXT NOT NULL UNIQUE,
	Secret				TEXT NOT NULL,
	BucketId			TEXT NOT NULL,
	BucketPrefix	TEXT,
	PRIMARY KEY(Uuid)
);

CREATE TABLE Team (
	Uuid						TEXT NOT NULL,
  Name						TEXT NOT NULL UNIQUE,
  Description			TEXT,
	CredentialUuid	TEXT NOT NULL UNIQUE,
	PRIMARY KEY(Uuid),
	FOREIGN KEY(CredentialUuid) REFERENCES AwsCredential(Uuid) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE Account (
	Uuid				TEXT NOT NULL,
	TeamUuid		TEXT,
	Username		TEXT NOT NULL,
	Password		TEXT NOT NULL,
  PRIMARY KEY(Uuid),
	FOREIGN KEY(TeamUuid) REFERENCES Team(Uuid) ON DELETE SET NULL ON UPDATE CASCADE
);

CREATE TABLE Mail (
	Uuid				TEXT NOT NULL,
	OwnerUuid		TEXT NOT NULL,
	PRIMARY KEY(Uuid)
);

CREATE TABLE Session (
	Uuid        TEXT,
	AccountUuid TEXT
);