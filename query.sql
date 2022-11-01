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

-- name: CreateAccount :one
INSERT INTO Account (
	uuid,
	username,
	password
) VALUES (
  ?, ?, ?
)
RETURNING *;