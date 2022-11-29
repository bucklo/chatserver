package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var dbPool *pgxpool.Pool
var ctx context.Context = context.Background()

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func connect(ctx context.Context) error {
	host := getEnv("POSTGRES_HOST", "localhost")
	port := getEnv("POSTGRES_PORT", "5432")
	user := getEnv("POSTGRES_USER", "postgres")
	password := getEnv("POSTGRES_PASSWORD", "postgres")

	connStr := fmt.Sprintf("postgres://%v:%v@%v:%v/chatserver?sslmode=disable", user, password, host, port)

	var err error

	for i := 1; i < 10; i++ {
		dbPool, err = pgxpool.Connect(ctx, connStr)
		if err != nil {
			fmt.Printf("Unable to connect to database. Retrying in %d seconds..\n", i*2)
			time.Sleep(time.Duration(i) * time.Second * 2)
			continue
		} else {
			log.Print("Successfully connected to DB")
			return err
		}
	}

	// Test the DB connection. This is only reached if the for-loop above failed
	// to establish a database connection.
	err = dbPool.Ping(ctx)
	return err
}

func InitializeDB() (*pgxpool.Pool, error) {
	err := connect(ctx)
	if err != nil {
		fmt.Print(err)
	}

	cmdTag, err := dbPool.Exec(ctx, "CREATE TABLE IF NOT EXISTS users(id SERIAL PRIMARY KEY, username VARCHAR(255) UNIQUE NOT NULL, password VARCHAR(255) NOT NULL, created_at TIMESTAMP NOT NULL DEFAULT NOW(), updated_at TIMESTAMP NOT NULL DEFAULT NOW())")
	if err != nil {
		fmt.Print(err)
	}

	log.Printf("%v, %v rows affected.", cmdTag, cmdTag.RowsAffected())

	return dbPool, err
}

func UserExists(username string) bool {
	var exists bool

	err := dbPool.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)", username).Scan(&exists)
	if err != nil {
		fmt.Print(err)
	}

	return exists
}

func AddUser(username, password string) {
	if UserExists(username) {
		log.Printf("User %v already exists", username)
		return
	}

	cmdTag, err := dbPool.Exec(ctx, "INSERT INTO users (username, password) VALUES ($1, $2)", username, password)
	if err != nil {
		fmt.Print(err)
	}

	log.Printf("%v, %v rows affected.", cmdTag, cmdTag.RowsAffected())
}
