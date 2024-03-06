package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
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
	fmt.Fprintf(w, "Welcome, %s!\n", name)

	for i := 0; i < numQuestions; i++ {
		operation := "+" // You can change the operation here if needed
		question, answer := generateQuestion(operation)

		// Print the question to the response
		fmt.Fprintf(w, "Question %d: %s\n", i+1, question)

		// Decode user answer from JSON request body
		var userAnswer struct {
			Answer int `json:"answer"`
		}
		if err := json.NewDecoder(r.Body).Decode(&userAnswer); err != nil {
			http.Error(w, "Failed to decode answer: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Check if user's answer matches the correct answer
		if userAnswer.Answer == answer {
			correctAnswers++
			fmt.Fprintf(w, "Correct!\n")
		} else {
			fmt.Fprintf(w, "Incorrect! The correct answer is %d.\n", answer)
		}
	}

	// Calculate score and send response
	score := correctAnswers * 100 / numQuestions
	fmt.Fprintf(w, "Your score is: %d", score)
}

func main() {
	// HTTP server setup
	http.HandleFunc("/mathques", handleMathQues)

	// Start the HTTP server
	fmt.Println("Server listening on port 8080")
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
