package AwsGateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type StorageActions struct {
	Download *manager.Downloader
	Upload   *manager.Uploader
}

func GetAwsCredential() (aws.Config, error) {
	User, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return aws.Config{}, err
	}
	return User, nil
}

// AWS S3 스토리지를 가져옵니다.
func GetStorage(User aws.Config) (*s3.Client, StorageActions) {
	Client := s3.NewFromConfig(User)
	Methods := StorageActions{
		Download: manager.NewDownloader(Client),
		Upload:   manager.NewUploader(Client),
	}

	return Client, Methods
}

// S3 스토리지에 있는 모든 파일을 가져옵니다.
func GetFilesInStorage(User aws.Config, Bucket string, BucketPrefix string, ProcessFunc func(StorageActions, string, string) error) error {
	Client, Methods := GetStorage(User)
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
			err = ProcessFunc(Methods, Bucket, aws.ToString(obj.Key))
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

func RemoveAllFiles() {
	// TODO: AWS S3 버킷 내부에 있는 모든 파일 삭제
	// GetFilesInStorage()을 한번 호출하고 본 함수를 호출하도록 유도
}

func IsVaildCredential() {
	// TODO: AWS Credential 유효성 검증 로직
	// 프로그램 시작시 일회성으로 DB에 있는 모든 Credential 검증하기
}
