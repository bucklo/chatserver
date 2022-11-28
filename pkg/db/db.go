package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func connect(ctx context.Context) (*pgxpool.Pool, error) {
	host := getEnv("POSTGRES_HOST", "localhost")
	port := getEnv("POSTGRES_PORT", "5432")
	user := getEnv("POSTGRES_USER", "postgres")
	password := getEnv("POSTGRES_PASSWORD", "postgres")

	connStr := fmt.Sprintf("postgres://%v:%v@%v:%v/chatserver?sslmode=disable", user, password, host, port)

	var dbPool *pgxpool.Pool
	var err error

	for i := 1; i < 10; i++ {
		dbPool, err = pgxpool.Connect(ctx, connStr)
		if err != nil {
			fmt.Printf("Unable to connect to database. Retrying in %d seconds..\n", i*2)
			time.Sleep(time.Duration(i) * time.Second * 2)
			continue
		} else {
			log.Print("Successfully connected to DB")
			return dbPool, err
		}
	}

	// Test the DB connection. This is only reached if the for-loop above failed
	// to establish a database connection.
	err = dbPool.Ping(ctx)
	return nil, err
}

func InitializeDB() (*pgxpool.Pool, error) {
	ctx := context.Background()

	dbPool, err := connect(ctx)
	if err != nil {
		fmt.Print(err)
	}

	cmdTag, err := dbPool.Exec(ctx, "CREATE TABLE IF NOT EXISTS users(id INT)")
	if err != nil {
		fmt.Print(err)
	}

	log.Printf("%v, %v rows affected.", cmdTag, cmdTag.RowsAffected())

	return dbPool, err
}
