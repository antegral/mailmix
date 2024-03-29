CREATE TABLE AwsCredential (
  Uuid          UUID NOT NULL,
  Region        TEXT NOT NULL,
  Id            TEXT NOT NULL UNIQUE,
  Secret        TEXT NOT NULL,
  BucketId      TEXT NOT NULL,
  BucketPrefix  TEXT,
  CreatedAt     TIMESTAMP NOT NULL DEFAULT NOW(),
  PRIMARY KEY(Uuid)
);

CREATE TABLE Team (
  Uuid            UUID NOT NULL,
  Name            TEXT NOT NULL UNIQUE,
  Description     TEXT,
  CredentialUuid  TEXT NOT NULL UNIQUE,
  CreatedAt       TIMESTAMP NOT NULL DEFAULT NOW(),
  PRIMARY KEY(Uuid),
  FOREIGN KEY(CredentialUuid) REFERENCES AwsCredential(Uuid) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE Account (
  Uuid        UUID NOT NULL,
  TeamUuid    UUID,
  Username    TEXT NOT NULL UNIQUE,
  Password    TEXT NOT NULL,
  MailAddress TEXT NOT NULL UNIQUE,
  IsQuit      BOOLEAN NOT NULL DEFAULT FALSE,
  CreatedAt   TIMESTAMP NOT NULL DEFAULT NOW(),
  PRIMARY KEY(Uuid),
  FOREIGN KEY(TeamUuid) REFERENCES Team(Uuid) ON DELETE SET NULL ON UPDATE CASCADE
);

CREATE TABLE MailBox (
  Uuid        UUID NOT NULL,
  Name        TEXT NOT NULL,
  OwnerUuid   UUID NOT NULL,
  Attributes  TEXT[] NOT NULL,
  Messages    BIGINT NOT NULL DEFAULT 0,
  Recent      BIGINT NOT NULL DEFAULT 0,
  Unseen      BIGINT NOT NULL DEFAULT 0,
  UidNext     BIGINT NOT NULL DEFAULT 0,
  UidValidity BIGINT NOT NULL DEFAULT 0,
  AppendLimit BIGINT NOT NULL DEFAULT 0,
  CreatedAt   TIMESTAMP NOT NULL DEFAULT NOW(),
  PRIMARY KEY(Uuid),
  FOREIGN KEY(OwnerUuid) REFERENCES Account(Uuid) ON DELETE CASCADE ON UPDATE CASCADE
);


CREATE TABLE Mail (
  Uid         BIGINT NOT NULL,
  Uuid        UUID NOT NULL,
  BoxUuid     UUID NOT NULL,
  Header      TEXT NOT NULL,
  SentFrom    TEXT NOT NULL,
  SentTo      TEXT NOT NULL,
  SentAt      TIMESTAMP NOT NULL,
  Hash        TEXT NOT NULL,
  Flags       VARCHAR(9)[] NOT NULL,
  Size        INTEGER NOT NULL,
  CreatedAt   TIMESTAMP NOT NULL DEFAULT NOW(),
  PRIMARY KEY(Uuid),
  FOREIGN KEY(TeamUuid) REFERENCES Team(Uuid) ON DELETE SET NULL ON UPDATE CASCADE
);

CREATE TABLE Attachment (
  Uuid        UUID NOT NULL,
  MailUuid    UUID NOT NULL,
	CID         TEXT NOT NULL,
	ContentType TEXT NOT NULL,
  CreatedAt   TIMESTAMP NOT NULL DEFAULT NOW(),
  PRIMARY KEY(Uuid)
);

CREATE TABLE EmbeddedFile (
  Uuid        UUID NOT NULL,
  MailUuid    UUID NOT NULL,
  ContentType TEXT NOT NULL,
  CreatedAt   TIMESTAMP NOT NULL DEFAULT NOW(),
  PRIMARY KEY(Uuid)
);

CREATE TABLE Session (
  Uuid        UUID NOT NULL,
  AccountUuid UUID NOT NULL,
  CreatedAt   TIMESTAMP NOT NULL DEFAULT NOW(),
  PRIMARY KEY(Uuid),
  FOREIGN KEY(AccountUuid) REFERENCES Account(Uuid) ON DELETE SET NULL ON UPDATE CASCADE
);
