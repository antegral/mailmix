package AwsGateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GetAwsCredential() (aws.Config, error) {
	User, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return aws.Config{}, err
	}
	return User, nil
}

// AWS S3 스토리지를 가져옵니다.
func GetStorage(User aws.Config) (*s3.Client, *manager.Downloader) {
	Client := s3.NewFromConfig(User)
	Manager := manager.NewDownloader(Client)

	return Client, Manager
}

// S3 스토리지에 있는 모든 파일을 다운받습니다.
func GetFilesInStorage(User aws.Config, Bucket string, BucketPrefix string, ProcessFunc func(*manager.Downloader, string, string) error) error {
	Client, Manager := GetStorage(User)
	var count int = 0

	// paginator를 만들고 S3 스토리지에 있는 모든 메일을 찾습니다.
	Paginator := s3.NewListObjectsV2Paginator(Client, &s3.ListObjectsV2Input{
		Bucket: &Bucket,
		Prefix: &BucketPrefix,
	})

	// 페이지를 넘겨가며 메일을 다운로드 합니다.
	var e error
	for Paginator.HasMorePages() {
		page, err := Paginator.NextPage(context.TODO())
		if err != nil {
			return err // paginator.NextPage() 예외 처리
		}
		for _, obj := range page.Contents {
			err = ProcessFunc(Manager, Bucket, aws.ToString(obj.Key))
			if err != nil {
				e = err
				break
			} else {
				count++
			}
		}
	}
	if e != nil {
		return e
	}
	return nil
}

func IsVaildCredential() {
	// TODO: AWS Credential 유효성 검증 로직
	// 프로그램 시작시 일회성으로 DB에 있는 모든 Credential 검증하기
}
