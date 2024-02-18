// main.go
package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/McCoyAle/arithmamom/db"
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
