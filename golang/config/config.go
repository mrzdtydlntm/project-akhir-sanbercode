package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB() *sqlx.DB {
	host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")
	user := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	dbname := os.Getenv("PGDBNAME")

	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Open("postgres", psqlUrl)
	if err != nil {
		log.Fatalf("Can't connect to database with err: %s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Database not responding with err: %s", err)
	}

	return db
}
