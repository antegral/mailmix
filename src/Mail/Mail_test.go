package Mail

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
)

func TestDownloadMail(t *testing.T) {
	Bucket := "ses-emails-antegral"
	BucketPrefix := ""

	if User, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"), config.WithSharedConfigProfile("dev-mailmix")); err != nil {
		t.Errorf(err.Error())
	} else {
		if err := DownloadMail(User, Bucket, BucketPrefix); err != nil {
			t.Errorf(err.Error())
		}
	}
}
