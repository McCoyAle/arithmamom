// main.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/McCoyAle/arithmamom/db"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/jackc/pgx/v5/pgxpool"
)

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


// connectToDB establishes a connection to the PostgreSQL database.
func connectToDB() (*pgxpool.Pool, error) {
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

	// Create a new database pool
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	fmt.Println("Connected to the database")
	return pool, nil
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
	// Connect to the PostgreSQL database
	conn, err := db.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Create user and score tables if they don't exist
	err = db.CreateUserTable(conn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.CreateScoreTable(conn)
	if err != nil {
		log.Fatal(err)
	}

	// Generate 10 random math questions and keep track of correct answers
	numQuestions := 10
	correctAnswers := 0
	for i := 0; i < numQuestions; i++ {
		operation := "+" // You can change the operation here if needed
		question, answer := generateQuestion(operation)
		userAnswer := getAnswerFromUser(question)
		if userAnswer == answer {
			fmt.Println("Correct!")
			correctAnswers++
		} else {
			fmt.Println("Incorrect!")
		}
	}

	// Insert score into the database
	err = db.InsertScore(conn, 1, correctAnswers) // Assuming userID is 1
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User and score inserted into the database successfully")
}
