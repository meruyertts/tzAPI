package db

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

// init initializes the database connection
func init() {
	log.Println("Initializing database connection...")
	DB = NewDB()
	log.Println("Database connection initialized successfully.")
}

// NewDB creates a new database connection and applies migrations
func NewDB() *sqlx.DB {
	err := godotenv.Load("conn.env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	connStr := os.Getenv("DATABASE_URL")
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting postgres:", err)
	}

	log.Println("Applying database migrations...")
	err = applyMigrations(db)
	if err != nil {
		log.Fatal("Error applying migrations:", err)

	}

	return db
}

func applyMigrations(db *sqlx.DB) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("error creating migrate instance: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver,
	)
	if err != nil {
		return fmt.Errorf("error creating migrate instance: %w", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error applying migrations: %w", err)
	}

	log.Println("Migrations applied successfully!")
	return nil
}
