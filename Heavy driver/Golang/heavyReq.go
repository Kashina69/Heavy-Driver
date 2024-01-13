package main

import (
	"fmt"
	"sync"
	"time"
)

// Placeholder for heavy computation
func heavyComputation(data map[string]interface{}) string {
	// Implement your heavy computation logic here
	fmt.Println("Running heavy computation:", data)
	time.Sleep(5 * time.Second) // Simulating heavy computation
	return "Heavy computation result"
}

// Function to run the heavy computation using a goroutine and channel
func runHeavyComputation(data map[string]interface{}, resultChan chan<- string) {
	go func() {
		result := heavyComputation(data)
		resultChan <- result
	}()
}

// Modify theHeavyReqProcessing function to use runHeavyComputation
func theHeavyReqProcessing(data map[string]interface{}, resultChan chan<- string) {
	// Modify this timeout if needed
	timeout := 10 * time.Minute
	theHeavyReqProcessingTimeout := time.After(timeout)

	go func() {
		select {
		case <-theHeavyReqProcessingTimeout:
			fmt.Println("Heavy computation timeout exceeded")
			resultChan <- "Heavy computation timeout exceeded"
		default:
			runHeavyComputation(data, resultChan)
		}
	}()
}

func main() {
	// Simulate running theHeavyReqProcessing
	data := map[string]interface{}{"key": "value"}
	resultChan := make(chan string, 1)

	theHeavyReqProcessing(data, resultChan)

	// Simulate receiving the result
	select {
	case result := <-resultChan:
		fmt.Println("Received result:", result)
		// Handle the result as needed
	}
}
