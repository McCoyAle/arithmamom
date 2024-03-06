package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
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
		num2 = rand.Intn(10) + 1
		answer = num1
		num1 = answer * num2
		question = fmt.Sprintf("%d / %d", num1, num2)
	default:
		panic("Invalid operation")
	}

	return question, answer
}

// handleMathQues handles the math Ques HTTP request.
func handleMathQues(w http.ResponseWriter, r *http.Request) {
	// Generate math Ques
	numQuestions := 10
	correctAnswers := 0
	for i := 0; i < numQuestions; i++ {
		operation := "+" // You can change the operation here if needed
		question, answer := generateQuestion(operation)

		// Print the question to the response
		fmt.Fprintf(w, "Question %d: %s\n", i+1, question)

		// Extract user answer from query parameter "answer"
		userAnswer := r.URL.Query().Get("answer")
		if userAnswer == "" {
			http.Error(w, "Missing answer parameter", http.StatusBadRequest)
			return
		}

		// Parse user answer to int
		userAnswerInt, err := strconv.Atoi(userAnswer)
		if err != nil {
			http.Error(w, "Invalid answer format", http.StatusBadRequest)
			return
		}

		// Check if user's answer matches the correct answer
		if userAnswerInt == answer {
			correctAnswers++
		}
	}

	// Calculate score and send response
	score := correctAnswers * 100 / numQuestions
	fmt.Fprintf(w, "Your score is: %d", score)
}

func main() {
	// HTTP server setup
	http.HandleFunc("/mathQues", handleMathQues)

	// Start the HTTP server
	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
