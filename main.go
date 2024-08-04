package main

// Import the required packages
import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// BalanceDataResponse struct
type BalanceDataResponse struct {
	Balance string `json:"balance"`
}

// Retrieve Balance data from Alephium API
func retrieveBalanceData(acct string) (float64, error) {
	// Get the Alephium Mainnet URL from the environment variable
	url := os.Getenv("ALEPHIUM_MAINNET_URL") + acct + "/balance"
	method := "GET"

	// Create a new HTTP client
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	// Check for errors
	if err != nil {
		return 0, err
	}

	// Send the request
	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	// Unmarshal the JSON response
	var r = new(BalanceDataResponse)
	err = json.Unmarshal(body, &r)
	if err != nil {
		return 0, err
	}

	// Convert the balance to a float
	balance, err := strconv.ParseFloat(r.Balance, 64)
	if err != nil {
		return 0, err
	}

	// Convert the balance to Alephium
	balance /= 1e18
	return balance, nil
}

func main() {
	// Load the .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the account balance from command line argument
	if len(os.Args) < 2 {
		fmt.Println("Please provide an account balance to query")
		return
	}
	acct := os.Args[1]

	// Retrieve the balance data
	balance, err := retrieveBalanceData(acct)
	fmt.Println("Balance: ", balance)
}
