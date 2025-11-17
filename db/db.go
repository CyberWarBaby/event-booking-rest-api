package db

import (
    "database/sql"
    "log"
    "os"

    _ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
    dsn := os.Getenv("DATABASE_URL")
    var err error
    DB, err = sql.Open("postgres", dsn)
    if err != nil {
        log.Fatalf("could not connect to database: %v", err)
    }

    if err := DB.Ping(); err != nil {
        log.Fatalf("could not ping database: %v", err)
    }

    DB.SetMaxOpenConns(10)
    DB.SetMaxIdleConns(5)

    createTables()
}

func createTables() {
    createUsersTable := `
    CREATE TABLE IF NOT EXISTS users(
        id SERIAL PRIMARY KEY,
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL
    );`
    _, err := DB.Exec(createUsersTable)
    if err != nil {
        log.Fatalf("could not create users table: %v", err)
    }

    createEventsTable := `
    CREATE TABLE IF NOT EXISTS events(
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        location TEXT NOT NULL,
        dateTime TIMESTAMP NOT NULL,
        user_id INTEGER REFERENCES users(id)
    );`
    _, err = DB.Exec(createEventsTable)
    if err != nil {
        log.Fatalf("could not create events table: %v", err)
    }

    createRegistrationsTable := `
    CREATE TABLE IF NOT EXISTS registrations(
        id SERIAL PRIMARY KEY,
        event_id INTEGER NOT NULL REFERENCES events(id),
        user_id INTEGER NOT NULL REFERENCES users(id)
    );`
    _, err = DB.Exec(createRegistrationsTable)
    if err != nil {
        log.Fatalf("could not create registrations table: %v", err)
    }
}
