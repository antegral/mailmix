package Database

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

func Init() error {
	_, err := GetDatabase()
	if err != nil {
		return err
	}

	return nil
}

func GetDatabase() (*sql.DB, error) {
	ConnString := os.Getenv("MAILMIX_DATABASE_CONNSTRING")

	db, err := sql.Open("pgx", ConnString)
	if err != nil {
		return nil, err
	}
	return db, err
}

func GetEnvFilePath() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Join(pwd, "/", ".env"), nil
}
