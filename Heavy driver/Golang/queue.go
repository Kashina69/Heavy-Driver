package main

import (
	"fmt"
	"sync"
	"time"
)

// Placeholder for heavy request processing
func heavyReqProcessing(userData map[string]interface{}) {
	// Implement your heavy request processing logic here
	fmt.Println("Processing heavy request:", userData)
	time.Sleep(2 * time.Second) // Simulating heavy processing
	fmt.Println("Heavy request processed")
}

var mu sync.Mutex
var userDataQueue = make([]map[string]interface{}, 0)

// Function to add user data to the queue
func addUserdataInQueue(currentRequestData map[string]interface{}) {
	mu.Lock()
	defer mu.Unlock()
	userDataQueue = append(userDataQueue, currentRequestData)
}

// Function to run the next task from the queue
func runNextTaskFromQueue() {
	mu.Lock()
	defer mu.Unlock()
	if len(userDataQueue) > 0 {
		heavyReqProcessing(userDataQueue[0])
		userDataQueue = userDataQueue[1:]
	}
}

func main() {
	// Simulate adding user data to the queue
	addUserdataInQueue(map[string]interface{}{"key1": "value1"})
	addUserdataInQueue(map[string]interface{}{"key2": "value2"})

	// Simulate running the next task from the queue
	runNextTaskFromQueue()
}
