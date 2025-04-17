package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Class-Connect-GRUPO-5/microservices-common/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

// DB stores the connection pool (pgxpool.Pool).
// It is used throughout the application to interact with the PostgreSQL database.
// The pool is created when the Connect function is called.
var DB *pgxpool.Pool

// Connect establishes a connection to the PostgreSQL database using environment variables.
// It also executes any necessary migrations to ensure required tables exist in the database.
func Connect() {
	// Retrieve database connection info from environment variables
	databaseInfo := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	logger.Logger.Debugf("Attempting to connect to database with connection string: %s", databaseInfo)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	logger.Logger.Debug("Creating database connection pool")
	DB, err = pgxpool.New(ctx, databaseInfo)

	if err != nil {
		logger.Logger.Errorf("Failed to create database connection pool: %v", err)
	}

	// Ping the database to ensure it's reachable
	err = DB.Ping(ctx)
	logger.Logger.Debug("Pinging the database to check connectivity")
	if err != nil {
		logger.Logger.Errorf("Database is not reachable: %v", err)
	}

	logger.Logger.Info("Database connection established successfully")

	// Run necessary migrations
	RunMigrations()
}

// RunMigrations executes the database migration process by reading SQL commands
// from a migration file and applying them to the database. It logs the progress
// and any errors encountered during the process.
//
// The function performs the following steps:
// 1. Reads the migration SQL file located at "./src/database/migrations.sql".
// 2. Executes the SQL commands in the migration file against the database.
// 3. Logs debug, error, and informational messages to indicate the status of the migration.
//
// If an error occurs while reading the migration file or executing the SQL commands,
// the error is logged, and the function continues execution.
func RunMigrations() {
	logger.Logger.Debug("Running database migrations")

	ctx := context.Background()

	migrationFilePath := os.Getenv("MIGRATION_FILE")
	if migrationFilePath == "" {
		logger.Logger.Debug("MIGRATION_FILE environment variable not set, using default path")
		migrationFilePath = "./src/database/migrations.sql"
	}

	migrationSQL, err := os.ReadFile(migrationFilePath)
	if err != nil {
		logger.Logger.Errorf("Error reading migration file: %v", err)
	}

	// Execute the SQL commands inside the migration file
	_, err = DB.Exec(ctx, string(migrationSQL))
	if err != nil {
		logger.Logger.Errorf("Error executing migrations: %v", err)
		return
	} else {
		logger.Logger.Info("Database migrations completed successfully")
	}
}
