package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Handle WebSocket connections
func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	fmt.Println("A user connected")

	// Infinite loop to read messages from the client
	for {
		var userData map[string]interface{}
		err := conn.ReadJSON(&userData)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("Received user data:", userData)

		// Process the request using your logic
		// Implement the logic in the takeReq function

		// Acknowledge that the request is received
		response := map[string]string{"message": "Request received successfully"}
		err = conn.WriteJSON(response)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections)

	// Start the server
	PORT := 3000
	fmt.Printf("Server is running on port %d\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}


// go get github.com/gorilla/websocket


// queue


