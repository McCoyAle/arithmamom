package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAnswerFromUser(t *testing.T) {
	// Create a new HTTP request with a question as a query parameter
	req, err := http.NewRequest("GET", "/get-answer?question=What is 2 + 2?", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the HTTP handler function
	handler := http.HandlerFunc(getAnswerFromUser)
	handler.ServeHTTP(rr, req)

	// Check the status code returned by the handler
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "User answer: 4"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
