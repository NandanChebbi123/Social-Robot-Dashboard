package main

// Structure of the Robot State

type RobotState struct {
	Status      string `json:"status"`
	Mood        string `json:"mood"`
	Activity    string `json:"activity"`
	PersonCount int    `json:"personCount"`
	State       string `json:"state"`
}

// Initial State of the Robot State, it is created by function.

func NewRobotState() RobotState {
	return RobotState{
		Status:      "online",
		Mood:        "neutral",
		Activity:    "waiting",
		PersonCount: 0,
		State:       "IDLE",
	}
}
