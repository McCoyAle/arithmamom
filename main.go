package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"net/http"
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
	// Retrieve user name from query parameter "name"
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Missing name parameter", http.StatusBadRequest)
		return
	}

	// Generate a random math question
	operation := "+"                           // You can change the operation here if needed
	question, _ := generateQuestion(operation) // Ignore the answer for now

	// Prepare the response buffer
	var responseBuffer bytes.Buffer

	// Write the welcome message to the response buffer
	responseBuffer.WriteString(fmt.Sprintf("Welcome, %s!\n", name))

	// Print the question to the response buffer
	responseBuffer.WriteString(fmt.Sprintf("Question: %s\n", question))

	// Write the accumulated response from the buffer to the response writer
	w.Write(responseBuffer.Bytes())
}

func main() {
	// Serve static files (HTML, CSS, JavaScript) from the "public" directory
	http.Handle("/", http.FileServer(http.Dir("public")))

	// Define HTTP endpoint for math questions and its handler
	http.HandleFunc("/mathques", handleMathQues)

	// Start the HTTP server
	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
