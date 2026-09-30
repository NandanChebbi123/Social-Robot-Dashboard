package main

// A single message in the conversation

type ConversationMessage struct {
	Speaker   string `json:"speaker"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// Stores the conversation history

type ConversationManager struct {
	Messages []ConversationMessage `json:"messages"`
}

// Creates an empty conversation.

func NewConversationManager() ConversationManager {
	return ConversationManager{
		Messages: []ConversationMessage{},
	}
}

// Adds a message to the conversation.

func (c *ConversationManager) AddMessage(message ConversationMessage) {
	c.Messages = append(c.Messages, message)
}

// Returns all conversation messages.

func (c *ConversationManager) GetMessages() []ConversationMessage {
	return c.Messages
}

// Clears the conversation.

func (c *ConversationManager) ClearMessages() {
	c.Messages = []ConversationMessage{}
}
