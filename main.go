package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/McCoyAle/arithmamom/db"
)

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
		num2 = rand.Intn(10) + 1
		answer = num1
		num1 = answer * num2
		question = fmt.Sprintf("%d / %d", num1, num2)
	default:
		panic("Invalid operation")
	}

	return question, answer
}

func getAnswerFromUser(question string) int {
	var userAnswer int
	fmt.Printf("What is %s? ", question)
	fmt.Scan(&userAnswer)
	return userAnswer
}

func handleMathQuiz(w http.ResponseWriter, r *http.Request) {
	// Generate math quiz
	numQuestions := 10
	correctAnswers := 0
	for i := 0; i < numQuestions; i++ {
		operation := "+"
		question, answer := generateQuestion(operation)
		userAnswer := getAnswerFromUser(question)
		if userAnswer == answer {
			correctAnswers++
		}
	}

	// Insert score into the database
	dbConn, err := db.ConnectToDB()
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to connect to the database: %v", err), http.StatusInternalServerError)
		return
	}
	defer dbConn.Close()

	err = db.InsertScore(dbConn, 1, correctAnswers) // Assuming userID is 1
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to insert score: %v", err), http.StatusInternalServerError)
		return
	}

	// Calculate score percentage
	score := correctAnswers * 100 / numQuestions

	// Respond with score
	fmt.Fprintf(w, "Your score is %d%%\n", score)
}

func main() {
	http.HandleFunc("/quiz", handleMathQuiz)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

	// Calculate score percentage
	score := correctAnswers * 100 / numQuestions

	// Respond with score
	fmt.Fprintf(w, "Your score is %d%%\n", score)
}

func main() {
	http.HandleFunc("/quiz", handleMathQuiz)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
