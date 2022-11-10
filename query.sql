-- name: CreateAwsCredential :one
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
RETURNING *;

-- name: CreateTeam :one
INSERT INTO Team (
	uuid,
	name,
	description,
	credentialuuid
) VALUES (
	?, ?, ?, ?
)
RETURNING *;

-- name: CreateAccount :one
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
RETURNING *;

-- name: CreateMailBox :one
INSERT INTO MailBox (
	uuid,
	name,
	owneruuid,
  attributes
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;

-- name: CreateMail :one
INSERT INTO Mail (
	uuid,
	boxuuid,
	header,
	sentfrom,
	sentto,
	sentat,
	content,
	flags,
	size
) VALUES (
	?, ?, ?, ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: CreateAttachment :one
INSERT INTO Attachment (
	uuid,
	mailuuid,
	cid,
	contenttype
) VALUES (
	?, ?, ?, ?
)
RETURNING *;

-- name: CreateEmbeddedFile :one
INSERT INTO EmbeddedFile (
	uuid,
	mailuuid,
	contenttype
) VALUES (
	?, ?, ?
)
RETURNING *;

-- name: CreateSession :one
INSERT INTO Session (
	uuid,
	accountuuid
) VALUES (
	?, ?
)
RETURNING *;

-- name: DeleteSessionByUser :exec
DELETE FROM Session
WHERE accountuuid = ?;

-- name: DeleteMailBox :exec
DELETE FROM MailBox
WHERE name = ?
AND owneruuid = ?;

-- name: DeleteAllMailInMailBox :exec
DELETE FROM Mail
WHERE boxuuid = ?;

-- name: GetAccountByUuid :one
SELECT * FROM Account
WHERE uuid = ?;

-- name: GetAccountByUsername :one
SELECT * FROM Account
WHERE username = ?;

-- name: GetAllMailBoxInfo :many
SELECT * FROM MailBox
WHERE owneruuid = ?;

-- name: GetOneMailBoxInfo :one
SELECT * FROM MailBox
WHERE name = ?
AND owneruuid = ?;

-- name: SetAccountPassword :exec
UPDATE Account
SET password = ?
WHERE username = ?;

-- name: RenameMailBox :exec
UPDATE MailBox
SET name = ?
WHERE name = ?
AND owneruuid = ?;

-- name: CountMailBox :one
SELECT count(*)
FROM MailBox
WHERE name = ?
AND owneruuid = ?;
