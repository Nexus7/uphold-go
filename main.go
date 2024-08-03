package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type BalanceDataResponse struct {
	Balance string `json:"balance"`
}

// Retrieve Balance data from Alephium API
func retrieveBalanceData(acct string) {
	url := "https://node.mainnet.alephium.org/addresses/" + acct + "/balance"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var r = new(BalanceDataResponse)
	err2 := json.Unmarshal(body, &r)
	if err2 != nil {
		log.Fatal(err2)
	}

	balance, err := strconv.ParseFloat(r.Balance, 64)
	if err != nil {
		log.Fatal(err)
	}

	balance /= 1e18
	fmt.Println("Balance: ", balance)
}

func main() {
	// Get the account balance from command line argument
	if len(os.Args) < 2 {
		fmt.Println("Please provide an account balance to query")
		return
	}
	acct := os.Args[1]

	retrieveBalanceData(acct)
}
