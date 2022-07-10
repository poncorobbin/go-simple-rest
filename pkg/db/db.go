package db

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Repo[T comparable] interface {
	Find() []T
	FindOne(string) T
	Save(T) T
}

func Connect() (*sql.DB, error) {
	dbUrl := os.Getenv("DB_URL")
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		return nil, err
	}

	return db, nil
}
