package Account

import (
	"crypto/sha512"
	b64 "encoding/base64"
)

type AwsCredential struct {
	Region     string
	ApiId      string
	ApiKey     string
	S3BucketId string
}

func GetB64Hash(plain string) string {
	sha := sha512.New()
	sha.Write([]byte("SALT+D3ntg902md+"))
	sha.Write([]byte(plain))
	b64hash := b64.StdEncoding.EncodeToString(sha.Sum(nil))
	return b64hash
}
