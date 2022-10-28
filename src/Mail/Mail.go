package Mail

import (
	"context"
	"fmt"
	"os"

	AwsGateway "antegral.net/mailmix/src/AwsGateway"
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
			// fmt.Print("DownloadMail > Key: ", Key, " => MailId: ", MailId, "\n")

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
			return nil
		}); e != nil {
		return e
	} else {
		return nil
	}
}

// func CategorizeMail(key string) {
// 	// TODO: Bucket에 들어온 메일들이 이메일 주소 별로 분류가 안되어 있을때 작동
// 	// 이메일을 파싱하여 Receiver에 적힌 이메일대로 분류 해야 함.
// 	ParsedString := strings.Split(key, "/")
// 	if len(ParsedString) > 1 {
// 		// fmt.Println("WARNING CategorizeMail > ")
// 	}
// }
