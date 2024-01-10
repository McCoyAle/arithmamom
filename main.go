package main

import (
	"fmt"
	"math/rand"
	"time"
)

// generateQuestion generates a random math question.
func generateQuestion() (string, int) {
	rand.Seed(time.Now().UnixNano())
	num1 := rand.Intn(10) + 1
	num2 := rand.Intn(10) + 1
	operator := []string{"+", "-", "*"}[rand.Intn(3)]
	question := fmt.Sprintf("%d %s %d", num1, operator, num2)
	var answer int

	switch operator {
	case "+":
		answer = num1 + num2
	case "-":
		answer = num1 - num2
	case "*":
		answer = num1 * num2
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
	fmt.Println("Welcome to Math Learning App!")

	score := 0
	numQuestions := 5

	for i := 0; i < numQuestions; i++ {
		question, correctAnswer := generateQuestion()
		userAnswer := getAnswerFromUser(question)

		if userAnswer == correctAnswer {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Printf("Wrong! The correct answer is %d.\n", correctAnswer)
		}
	}

	fmt.Printf("\nGame Over! Your score is %d/%d.\n", score, numQuestions)
}
