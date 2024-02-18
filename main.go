// main.go

package main

import (
	"fmt"
	"log"

	"github.com/McCoyAle/arithmamom/database" // Adjust the import path accordingly
)

// User and Score structs

func main() {
	// Connect to the PostgreSQL database using environment variables
	conn, err := database.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Create user and score tables if they don't exist
	err = database.CreateUserTable(conn)
	if err != nil {
		log.Fatal(err)
	}

	err = database.CreateScoreTable(conn)
	if err != nil {
		log.Fatal(err)
	}

	// Example: Insert a new user and score into the database
	err = database.InsertUser(conn, "testuser", "testpassword")
	if err != nil {
		log.Fatal(err)
	}

	err = database.InsertScore(conn, 1, 3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User and score inserted into the database successfully")
}
