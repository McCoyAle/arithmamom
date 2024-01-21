package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// User represents the user model.
type User struct {
	ID       int
	Username string
	Password string
}

// Score represents the math score model.
type Score struct {
	ID        int
	UserID    int
	Score     int
	Timestamp time.Time
}

// generateQuestion generates a random math question.
func generateQuestion(operation string) (string, int) {
	rand.Seed(time.Now().UnixNano())
	num1 := rand.Intn(50) + 1
	num2 := rand.Intn(50) + 1
	var question string
	var answer int

	switch operation {
	case "+":
		question = fmt.Sprintf("%d + %d", num1, num2)
		answer = num1 + num2
	case "-":
		question = fmt.Sprintf("%d - %d", num1, num2)
		answer = num1 - num2
	case "*":
		num1 = rand.Intn(10) + 1
		num2 = rand.Intn(10) + 1
		question = fmt.Sprintf("%d * %d", num1, num2)
		answer = num1 * num2
	case "/":
		// Ensure a whole number division for simplicity
		num2 = rand.Intn(10) + 1
		answer = num1
		num1 = answer * num2
		question = fmt.Sprintf("%d / %d", num1, num2)
	default:
		panic("Invalid operation")
	}

	return question, answer
}

// getAnswerFromUser gets the user's answer for the math question.
func getAnswerFromUser(question string) int {
	var userAnswer int
	fmt.Printf("What is %s? ", question)
	fmt.Scan(&userAnswer)
	return userAnswer
}

// connectToDB establishes a connection to the PostgreSQL database using environment variables.
func connectToDB() (*pgxpool.Pool, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	conn, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	fmt.Println("Connected to the database")
	return conn, nil
}

// createUserTable creates the user table in the database.
func createUserTable(conn *pgxpool.Pool) error {
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

// createScoreTable creates the score table in the database.
func createScoreTable(conn *pgxpool.Pool) error {
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

// insertUser inserts a new user into the database.
func insertUser(conn *pgxpool.Pool, user User) error {
	_, err := conn.Exec(context.Background(), "INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, user.Password)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	return nil
}

// insertScore inserts a new score into the database.
func insertScore(conn *pgxpool.Pool, score Score) error {
	_, err := conn.Exec(context.Background(), "INSERT INTO scores (user_id, score) VALUES ($1, $2)", score.UserID, score.Score)
	if err != nil {
		return fmt.Errorf("failed to insert score: %w", err)
	}

	return nil
}

func main() {
	// Connect to the PostgreSQL database using environment variables
	conn, err := connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Create user and score tables if they don't exist
	err = createUserTable(conn)
	if err != nil {
		log.Fatal(err)
	}

	err = createScoreTable(conn)
	if err != nil {
		log.Fatal(err)
	}

	// Example: Insert a new user and score into the database
	user := User{Username: "testuser", Password: "testpassword"}
	err = insertUser(conn, user)
	if err != nil {
		log.Fatal(err)
	}

	score := Score{UserID: 1, Score: 3}
	err = insertScore(conn, score)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User and score inserted into the database successfully")
}
