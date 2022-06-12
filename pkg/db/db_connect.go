package db

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Repo interface {
	Find() interface{}
	FindOne(id string) interface{}
	Save(interface{}) interface{}
}

func Connect() (*sql.DB, error) {
	dbUrl := os.Getenv("DB_URL")
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		return nil, err
	}

	return db, nil
}
