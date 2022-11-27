package db

import "database/sql"

func Connect() *sql.DB {
	// Connect to database
	connStr := "postgres://postgres:postgres@localhost:5432/chatserver?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}

	return db
}
