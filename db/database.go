// db/database.go
package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
)

// User represents the user model.
type User struct {
	ID       int
	Username string
	Password string
	Name     string
	Email    string
}

// Score represents the math score model.
type Score struct {
	ID        int
	UserID    int
	Score     int
	Timestamp time.Time
}

// ConnectToDB establishes a connection to the PostgreSQL database.
func ConnectToDB() (*sql.DB, error) {
	// Retrieve database credentials from AWS Secrets Manager
	secretName := "arithmamom-app"
	region := "us-east-1"

	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	// Create Secrets Manager client
	svc := secretsmanager.NewFromConfig(config)

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve secret value: %w", err)
	}

	// Parse the secret string to extract database credentials
	var secretString string
	if result.SecretString != nil {
		secretString = *result.SecretString
	} else {
		return nil, fmt.Errorf("secret string is empty")
	}

	// Parse secret string to extract database credentials
	var dbCredentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     string `json:"port"`
		Database string `json:"database"`
	}

	err = json.Unmarshal([]byte(secretString), &dbCredentials)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database credentials: %w", err)
	}

	// Construct the connection string
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", dbCredentials.Username, dbCredentials.Password, dbCredentials.Host, dbCredentials.Port, dbCredentials.Database)

	// Create a new database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	log.Println("Connected to the database")
	return db, nil
}

// CreateUserTable creates the user table in the database.
func CreateUserTable(conn *pgxpool.Pool) error {
	_, err := conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) UNIQUE NOT NULL,
			password VARCHAR(50) NOT NULL
		)
	`)

	if err != nil {
		return fmt.Errorf("failed to create user table: %w", err)
	}

	return nil
}

// CreateScoreTable creates the score table in the database.
func CreateScoreTable(conn *pgxpool.Pool) error {
	_, err := conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS scores (
			id SERIAL PRIMARY KEY,
			user_id INTEGER REFERENCES users(id),
			score INTEGER NOT NULL,
			timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)

	if err != nil {
		return fmt.Errorf("failed to create score table: %w", err)
	}

	return nil
}

// InsertUser inserts a new user into the database.
func InsertUser(conn *pgxpool.Pool, username, password string) error {
	_, err := conn.Exec(context.Background(), "INSERT INTO users (username, password) VALUES ($1, $2)", username, password)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	return nil
}

// InsertScore inserts a new score into the database.
func InsertScore(conn *pgxpool.Pool, userID, score int) error {
	_, err := conn.Exec(context.Background(), "INSERT INTO scores (user_id, score) VALUES ($1, $2)", userID, score)
	if err != nil {
		return fmt.Errorf("failed to insert score: %w", err)
	}

	return nil
}

