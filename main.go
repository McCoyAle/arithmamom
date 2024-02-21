package main

import (
	"fmt"
	"log"
	"math/rand"
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

func main() {
	var userName string
	fmt.Print("Please tell me your name: ")
	fmt.Scanln(&userName)
	fmt.Printf("Welcome, %s, to Math with Addi's Mama!\n", userName)

	conn, err := db.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	err = db.CreateUserTable(conn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.CreateScoreTable(conn)
	if err != nil {
		log.Fatal(err)
	}

	numQuestions := 10
	correctAnswers := 0
	for i := 0; i < numQuestions; i++ {
		operation := "+"
		question, answer := generateQuestion(operation)
		userAnswer := getAnswerFromUser(question)
		if userAnswer == answer {
			fmt.Printf("That's right %s!\n", userName)
			correctAnswers++
		} else {
			fmt.Printf("Sorry! The correct answer is %d.\n", answer)
		}
	}

	// Assuming userID is 1. You might want to dynamically retrieve it based on user information.
	userID := 1
	err = db.InsertScore(conn, userID, correctAnswers)
	if err != nil {
		log.Fatal(err)
	}

	score := correctAnswers * 100 / numQuestions
	fmt.Printf("\nSorry, %s! The Game is Over. Your score is %d/%d.\n", userName, score, numQuestions)
}
