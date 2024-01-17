package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/lib/pq"
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

// connectToDB establishes a connection to the PostgreSQL database.
func connectToDB() (*sql.DB, error) {
	connStr := "user=username dbname=math_scores sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to the database")
	return db, nil
}

// createUserTable creates the user table in the database.
func createUserTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) UNIQUE NOT NULL,
			password VARCHAR(50) NOT NULL
		)
	`)

	return err
}

// createScoreTable creates the score table in the database.
func createScoreTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS scores (
			id SERIAL PRIMARY KEY,
			user_id INTEGER REFERENCES users(id),
			score INTEGER NOT NULL,
			timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)

	return err
}

// insertUser inserts a new user into the database.
func insertUser(db *sql.DB, user User) error {
	_, err := db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, user.Password)
	return err
}

// insertScore inserts a new score into the database.
func insertScore(db *sql.DB, score Score) error {
	_, err := db.Exec("INSERT INTO scores (user_id, score) VALUES ($1, $2)", score.UserID, score.Score)
	return err
}

func main() {
	// Connect to the PostgreSQL database
	db, err := connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create user and score tables if they don't exist
	err = createUserTable(db)
	if err != nil {
		log.Fatal(err)
	}

	err = createScoreTable(db)
	if err != nil {
		log.Fatal(err)
	}

	// Example: Insert a new user and score into the database
	user := User{Username: "testuser", Password: "testpassword"}
	err = insertUser(db, user)
	if err != nil {
		log.Fatal(err)
	}

	score := Score{UserID: 1, Score: 3}
	err = insertScore(db, score)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User and score inserted into the database successfully")
}
