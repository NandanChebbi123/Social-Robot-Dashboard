package main

import (
	"fmt"
	"net/http"
)

func main() {

	// Create the robot's initial state
	robot := NewRobotState()

	// Create the state machine
	stateMachine := NewStateMachine(&robot)

	// Create the conversation manager
	conversations := NewConversationManager()

	// Create the API
	api := NewAPI(
		&robot,
		&stateMachine,
		&conversations,
	)

	// Create HTTP router
	mux := http.NewServeMux()

	// Register API routes
	api.RegisterRoutes(mux)

	// Start HTTP server
	fmt.Println("Social Robot Dashboard backend running on http://localhost:8080")

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Println("Server error:", err)
	}
}
