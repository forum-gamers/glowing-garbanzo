package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	h "github.com/forum-gamers/glowing-garbanzo/helpers"
)

func Conn() {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		url = "user=apple password= dbname=forum-gamers-community sslmode=disable"
	}
	db, err := sql.Open("postgres", url)
	h.PanicIfError(err)

	DB = db
	log.Println("connected to the database")
}
