package Account

import (
	"crypto/sha512"
	b64 "encoding/base64"
)

type AwsCredential struct {
	Region       string
	ApiId        string
	ApiKey       string
	BucketId     string
	BucketPrefix string
}

func GetB64Hash(Plain string) string {
	sha := sha512.New()
	sha.Write([]byte("SALT+D3ntg902md+"))
	sha.Write([]byte(Plain))
	b64hash := b64.StdEncoding.EncodeToString(sha.Sum(nil))
	return b64hash
}
