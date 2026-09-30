package main

type RobotEvent string

const (
	PersonDetected   RobotEvent = "PERSON_DETECTED"
	PersonLeft       RobotEvent = "PERSON_LEFT"
	GreetingFinished RobotEvent = "GREETING_FINISHED"
	SpeechReceived   RobotEvent = "SPEECH_RECEIVED"
	ResponseReady    RobotEvent = "RESPONSE_READY"
	ResponseFinished RobotEvent = "RESPONSE_FINISHED"
)
