package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex
var oldUserData map[string]interface{}

// Placeholder for heavy request processing
func heavyReqProcessing(userData map[string]interface{}, resultChan chan<- string) {
	// Implement your heavy request processing logic here
	fmt.Println("Processing heavy request:", userData)
	time.Sleep(5 * time.Second) // Simulating heavy processing
	resultChan <- "Heavy request processed"
}

// Function to run the heavy computation using a goroutine and channel
func runHeavyReqProcessing(userData map[string]interface{}, resultChan chan<- string) {
	go func() {
		heavyReqProcessing(userData, resultChan)
	}()
}

// Middleware to handle user requests
func takeReq(userData map[string]interface{}) {
	mu.Lock()
	defer mu.Unlock()

	// Your logic to check if the request is duplicate
	if compareJSON(userData, oldUserData) {
		return
	}

	// If you want to clear old request
	// Note: Timeout cancellation is a bit different in Go
	// You may want to use context.Context for a more robust solution
	oldUserData = nil

	// Create a channel to receive the result from heavyReqProcessing
	resultChan := make(chan string, 1)

	// Call the function for heavy request processing
	runHeavyReqProcessing(userData, resultChan)

	// Wait for the result or a timeout
	timeout := 10 * time.Second
	select {
	case result := <-resultChan:
		fmt.Println("Received result:", result)
		// Handle the result as needed
	case <-time.After(timeout):
		fmt.Println("Heavy request processing timeout exceeded")
	}
}

// Function to compare JSON objects
func compareJSON(a, b map[string]interface{}) bool {
	aJSON, errA := json.Marshal(a)
	bJSON, errB := json.Marshal(b)

	if errA != nil || errB != nil {
		return false
	}

	return string(aJSON) == string(bJSON)
}

func main() {
	// Simulate running takeReq
	data := map[string]interface{}{"key": "value"}
	takeReq(data)
}
