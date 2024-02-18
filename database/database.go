// database/database.go

package database

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// ConnectToDB establishes a connection to the PostgreSQL database using environment variables.
func ConnectToDB() (*pgxpool.Pool, error) {
	// Database connection setup code
}

// CreateUserTable creates the user table in the database.
func CreateUserTable(conn *pgxpool.Pool) error {
	// User table creation SQL
}

// CreateScoreTable creates the score table in the database.
func CreateScoreTable(conn *pgxpool.Pool) error {
	// Score table creation SQL
}

// InsertUser inserts a new user into the database.
func InsertUser(conn *pgxpool.Pool, username, password string) error {
	// User insertion SQL
}

// InsertScore inserts a new score into the database.
func InsertScore(conn *pgxpool.Pool, userID, score int) error {
	// Score insertion SQL
}
