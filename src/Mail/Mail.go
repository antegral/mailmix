package Mail

import (
	"context"
	"fmt"
	"io"
	"os"

	AwsGateway "antegral.net/mailmix/src/AwsGateway"
	"github.com/DusanKasan/parsemail"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

// 메일이 저장될 파일을 만들고 다운로드 합니다.
func DownloadMail(User aws.Config, Bucket string, BucketPrefix string) error {
	if e := AwsGateway.GetFilesInStorage(User, Bucket, BucketPrefix,
		func(Action AwsGateway.StorageActions, Bucket string, Key string) error {
			// 메일 고유번호 설정
			MailId := uuid.New()

			// 현재 경로 가져오기
			pwd, err := os.Getwd()
			if err != nil {
				return err // os.Getwd() 예외처리
			}

			// 경로 포맷팅
			FolderPath := fmt.Sprint(pwd, "/MailStorage/")
			FilePath := fmt.Sprint(pwd, "/MailStorage/", MailId)

			// os.Create로 파일 생성
			os.MkdirAll(FolderPath, os.ModePerm)
			File, err := os.Create(FilePath)
			if err != nil {
				return err // os.Create() 예외처리
			}

			// 다운이 끝나면 알아서 파일 닫기
			defer File.Close()

			// 메일 다운로드
			if _, err = Action.Download.Download(context.TODO(), File, &s3.GetObjectInput{
				Bucket: &Bucket,
				Key:    &Key,
			}); err != nil {
				return err
			}

			if err := IndexMail(File); err != nil {
				return err
			}
			return nil
		}); e != nil {
		return e
	}
	return nil
}

func IndexMail(Mail io.Reader) error {
	_, err := parsemail.Parse(Mail)
	if err != nil {
		return err
	}
	return nil
}
