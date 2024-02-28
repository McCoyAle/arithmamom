package db_test

import (
	"testing"

	"github.com/McCoyAle/arithmamom/db"
	"github.com/stretchr/testify/assert"
)

// TestConnectToDB tests the ConnectToDB function.
func TestConnectToDB(t *testing.T) {
	pool, err := db.ConnectToDB()
	assert.NoError(t, err)
	assert.NotNil(t, pool)

	// Close the pool
	pool.Close()
}

// TestCreateUserTable tests the CreateUserTable function.
func TestCreateUserTable(t *testing.T) {
	// Establish a connection to the database
	pool, err := db.ConnectToDB()
	assert.NoError(t, err)
	defer pool.Close()

	// Create user table
	err = db.CreateUserTable(pool)
	assert.NoError(t, err)

	// TODO: Optionally, add additional assertions to verify the table creation
}

// TestCreateScoreTable tests the CreateScoreTable function.
func TestCreateScoreTable(t *testing.T) {
	// Establish a connection to the database
	pool, err := db.ConnectToDB()
	assert.NoError(t, err)
	defer pool.Close()

	// Create score table
	err = db.CreateScoreTable(pool)
	assert.NoError(t, err)

	// TODO: Optionally, add additional assertions to verify the table creation
}

// Add more test functions at a later time
