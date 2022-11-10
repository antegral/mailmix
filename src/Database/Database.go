package Database

import (
	"database/sql"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
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

	db, err := sql.Open("postgres", ConnString)
	if err != nil {
		return nil, err
	}
	return db, err
}
