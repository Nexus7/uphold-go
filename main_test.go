package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		os.Exit(1)
	}

	// Run tests
	code := m.Run()

	// Clean up resources if needed

	os.Exit(code)
}

func TestValidAlephiumAddress(t *testing.T) {
	validAddress := "19WzSnmNC1SQ6v7RpFFXhpcMcFSiwM4nKTSdbwgSJfSHy"
	balance, err := retrieveBalanceData(validAddress)

	if err != nil {
		t.Errorf("Error retrieving balance data: %s", err)
	}

	if balance <= 0 {
		t.Errorf("Expected balance to be greater than 0, got %f", balance)
	}
}
