package postgres

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func connString() string {
	format := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	return fmt.Sprintf(
		format,
		host, user, password, dbname, port,
	)
}

func Connect() error {
	var err error

	db, err = sqlx.Connect("postgres", connString())
	if err != nil {
		log.Println("error connecting to database", err)
		return err
	}

	return nil
}

func DB() *sqlx.DB {
	return db
}
