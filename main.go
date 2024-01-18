package main

import (
	"fmt"
	"math/rand"
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
	fmt.Print("Please tell me your name: ")
	var userName string
	fmt.Scanln(&userName)

	fmt.Printf("Welcome, %s, to Math with Addi's Mama!\n", userName)

	// Code base for choosing a math operation to perform
	fmt.Print("Choose a mathematical operation (+, -, *, /): ")
	var selectedOperation string
	fmt.Scan(&selectedOperation)

	score := 0
	numQuestions := 5

	for i := 0; i < numQuestions; i++ {
		question, correctAnswer := generateQuestion(selectedOperation)
		userAnswer := getAnswerFromUser(question)

		if userAnswer == correctAnswer {
			fmt.Printf("That's right, %s!\n", userName)
			score++
		} else {
			fmt.Printf("The correct answer is %d.\n", correctAnswer)
		}
	}

	percentage := float64(score) / float64(numQuestions) * 100
	fmt.Printf("\nSorry, %s! The Game is Over. Your score is %.2f%%.\n", userName, percentage)
}
