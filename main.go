package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// Initialize a global random generator
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// generateQuestion generates a random math question.
func generateQuestion(operation string) (string, int) {
	num1 := r.Intn(50) + 1
	num2 := r.Intn(50) + 1
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
		num1 = r.Intn(10) + 1
		num2 = r.Intn(10) + 1
		question = fmt.Sprintf("%d * %d", num1, num2)
		answer = num1 * num2
	case "/":
		num2 = r.Intn(10) + 1
		answer = num1
		num1 = answer * num2
		question = fmt.Sprintf("%d / %d", num1, num2)
	default:
		panic("Invalid operation")
	}

	return question, answer
}

// handleMathQues handles the math ques HTTP request.
func handleMathQues(w http.ResponseWriter, r *http.Request) {
	// Generate math ques
	numQuestions := 10
	correctAnswers := 0

	// Retrieve user name from query parameter "name"
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Missing name parameter", http.StatusBadRequest)
		return
	}

	var responseBuffer bytes.Buffer // Create a buffer to accumulate the response

	// Write the welcome message to the response buffer
	responseBuffer.WriteString(fmt.Sprintf("Welcome, %s!\n", name))

	for i := 0; i < numQuestions; i++ {
		operation := "+" // You can change the operation here if needed
		question, answer := generateQuestion(operation)

		// Print the question to the response buffer
		responseBuffer.WriteString(fmt.Sprintf("Question %d: %s\n", i+1, question))

		// Decode user answer from query parameter "answer"
		userAnswerStr := r.URL.Query().Get("answer")
		userAnswer, err := strconv.Atoi(userAnswerStr)
		if err != nil {
			http.Error(w, "Failed to parse answer: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Check if user's answer matches the correct answer
		if userAnswer == answer {
			correctAnswers++
			responseBuffer.WriteString("Correct!\n")
		} else {
			responseBuffer.WriteString(fmt.Sprintf("Incorrect! The correct answer is %d.\n", answer))
		}
	}

	// Calculate score and append it to the response buffer
	score := correctAnswers * 100 / numQuestions
	responseBuffer.WriteString(fmt.Sprintf("Your score is: %d", score))

	// Write the accumulated response from the buffer to the response writer
	w.Write(responseBuffer.Bytes())
}

func main() {
	// Serve static files (HTML, CSS, JavaScript) from the "static" directory
	http.Handle("/", http.FileServer(http.Dir("public")))

	// Define HTTP endpoints and handlers
	http.HandleFunc("/mathques", handleMathQues)

	// Start the HTTP server
	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
