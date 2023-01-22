package Mail

import (
	"context"
	"fmt"
	"io"
	"os"

	AwsGateway "antegral.net/mailmix/src/AwsGateway"
	"antegral.net/mailmix/src/Database"
	ORM "antegral.net/mailmix/src/Database/Sqlc"
	"github.com/DusanKasan/parsemail"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/zeebo/blake3"
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
	// TODO: 메일 저장 후 DB에 기록

	ParsedMail, err := parsemail.Parse(Mail)
	if err != nil {
		return err
	}

	ctx := context.Background()
	database, err := Database.GetDatabase()
	if err != nil {
		return nil
	}

	// Import ORM
	Queries := ORM.New(database)

	h := blake3.New()
	h.Write([]byte(Mail))
	Hash := h.Sum(nil)

	Queries.CreateMail(ctx, ORM.CreateMailParams{
		Uuid: uuid.New(),
		Boxuuid: uuid.New(), // TODO: User의 기본 Mailbox UUID 가져와서 넣기
		Header:   ParsedMail.Header,
		Sentfrom: ParsedMail.Sender.String(),
		Sentto:   ParsedMail.To[0],
		Sentat:   ParsedMail.Date,
		hash:  Hash,
		Flags:    string
		Size:     int32
	})

	return nil
}
