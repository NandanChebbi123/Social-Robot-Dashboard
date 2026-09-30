package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type API struct {
	Robot         *RobotState
	StateMachine  *StateMachine
	Conversations *ConversationManager

	Mutex sync.Mutex
}

func NewAPI(
	robot *RobotState,
	stateMachine *StateMachine,
	conversations *ConversationManager,
) *API {
	return &API{
		Robot:         robot,
		StateMachine:  stateMachine,
		Conversations: conversations,
	}
}

// GET /api/robot
func (api *API) GetRobot(w http.ResponseWriter, r *http.Request) {

	api.Mutex.Lock()
	defer api.Mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(api.Robot)
}

// Data received when an event is sent.
type EventRequest struct {
	Event string `json:"event"`
}

// POST /api/events
func (api *API) HandleEvent(w http.ResponseWriter, r *http.Request) {

	var request EventRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	api.Mutex.Lock()
	defer api.Mutex.Unlock()

	event := RobotEvent(request.Event)

	api.StateMachine.HandleEvent(event)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(api.Robot)
}

// GET /api/conversations
// POST /api/conversations
func (api *API) HandleConversations(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:

		api.Mutex.Lock()
		defer api.Mutex.Unlock()

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(api.Conversations.GetMessages())

	case http.MethodPost:

		var message ConversationMessage

		err := json.NewDecoder(r.Body).Decode(&message)

		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		api.Mutex.Lock()
		defer api.Mutex.Unlock()

		api.Conversations.AddMessage(message)

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(message)

	default:

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Registers all API routes.
func (api *API) RegisterRoutes(mux *http.ServeMux) {

	mux.HandleFunc("/api/robot", api.GetRobot)

	mux.HandleFunc("/api/events", api.HandleEvent)

	mux.HandleFunc("/api/conversations", api.HandleConversations)
}
