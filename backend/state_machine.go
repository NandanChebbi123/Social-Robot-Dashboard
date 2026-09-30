package main

type StateMachine struct {
	Robot *RobotState
}

func NewStateMachine(robot *RobotState) StateMachine {
	return StateMachine{
		Robot: robot,
	}
}

func (sm *StateMachine) HandleEvent(event RobotEvent) {

	switch sm.Robot.State {

	case "IDLE":

		if event == PersonDetected {
			sm.Robot.State = "GREETING"
			sm.Robot.Activity = "greeting visitor"
			sm.Robot.Mood = "happy"
			sm.Robot.PersonCount = 1
		}

	case "GREETING":

		if event == GreetingFinished {
			sm.Robot.State = "LISTENING"
			sm.Robot.Activity = "listening"
			sm.Robot.Mood = "neutral"
		}

	case "LISTENING":

		if event == SpeechReceived {
			sm.Robot.State = "THINKING"
			sm.Robot.Activity = "thinking"
		}

		if event == PersonLeft {
			sm.Robot.State = "IDLE"
			sm.Robot.Activity = "waiting"
			sm.Robot.Mood = "neutral"
			sm.Robot.PersonCount = 0
		}

	case "THINKING":

		if event == ResponseReady {
			sm.Robot.State = "TALKING"
			sm.Robot.Activity = "talking"
		}

	case "TALKING":

		if event == ResponseFinished {
			sm.Robot.State = "LISTENING"
			sm.Robot.Activity = "listening"
			sm.Robot.Mood = "neutral"
		}
	}
}
