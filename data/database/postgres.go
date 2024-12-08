package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // for connecting sqlx to postgres
)

const (
	ErrCantConnectToDB = "error on connecting to database: %w"
	ErrCantPingDB      = "cant ping to database: %w"
	MsgConnectedToDB   = "connected to database successfully"
	MsgDBNotYetReady   = "Postgres not yet ready..."

	maxRetry = 10
)

var (
	retryCount int
)

func New() *sqlx.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	for {
		db, err := sqlx.Open("postgres", dsn)
		if err != nil {
			log.Println(MsgDBNotYetReady)
			retryCount++
		} else {
			if err := db.Ping(); err != nil {
				log.Println(MsgDBNotYetReady)
				retryCount++
			} else {
				log.Println(MsgConnectedToDB)
				return db
			}
		}

		if retryCount > maxRetry {
			panic(fmt.Errorf(ErrCantConnectToDB, err))
		}

		log.Println("Backing off for two seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}
