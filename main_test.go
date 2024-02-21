// main_test.go

package main

import (
	"os"
	"testing"
)

func TestGenerateQuestion(t *testing.T) {
	// Test addition
	question, answer := generateQuestion("+")
	if question == "" || answer == 0 {
		t.Errorf("Expected a non-empty question and non-zero answer for addition")
	}

	// Test subtraction
	question, answer = generateQuestion("-")
	if question == "" || answer == 0 {
		t.Errorf("Expected a non-empty question and non-zero answer for subtraction")
	}

	// Test multiplication
	question, answer = generateQuestion("*")
	if question == "" || answer == 0 {
		t.Errorf("Expected a non-empty question and non-zero answer for multiplication")
	}

	// Test division
	question, answer = generateQuestion("/")
	if question == "" || answer == 0 {
		t.Errorf("Expected a non-empty question and non-zero answer for division")
	}

	// Test invalid operation
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected generateQuestion to panic for an invalid operation")
		}
	}()
	generateQuestion("invalid")
}

func TestGetAnswerFromUser(t *testing.T) {
	// Simulate user input for testing
	input := "42\n" // Assuming the user enters 42
	expected := 42

	// Mock standard input
	stdin := setupStdin(input)
	defer stdin.Close()

	// Test getAnswerFromUser function
	userAnswer := getAnswerFromUser("What is the answer?")
	if userAnswer != expected {
		t.Errorf("Expected user answer to be %d, but got %d", expected, userAnswer)
	}
}

// Mock standard input for testing
func setupStdin(input string) *os.File {
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.Write([]byte(input))
	w.Close()
	return r
}
