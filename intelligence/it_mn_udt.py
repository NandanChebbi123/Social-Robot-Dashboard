from pydantic import BaseModel, Field
from typing import Optional

class Entity(BaseModel):
    name: str
    entity_type: str

class Intent(BaseModel):
    name: str
    confidence: float

class Meaning(BaseModel):
    topic: Optional[str] = None
    action: Optional[str] = None
    subject: Optional[str] = None
    object: Optional[str] = None


class UnderstandingResult(BaseModel):

    session_id: str
    original_input: str
    intent: Intent

    entities: list[Entity] = Field(
        default_factory=list
    )

    meaning: Meaning

    requires_clarification: bool = False

SUPPORTED_INTENTS = {
    "greeting",
    "farewell",
    "information_request",
    "question",
    "command",
    "statement",
    "opinion",
    "confirmation",
    "denial",
    "memory_request",
    "unknown"
}

def detect_intent(text: str) -> Intent:

    text_lower = text.lower().strip()

    greetings = [
        "hello",
        "hi",
        "hey",
        "good morning",
        "good afternoon",
        "good evening"
    ]

    if any(
        text_lower == greeting
        or text_lower.startswith(greeting + " ")
        for greeting in greetings
    ):
        return Intent(
            name="greeting",
            confidence=0.95
        )

    farewells = [
        "bye",
        "goodbye",
        "see you",
        "good night"
    ]

    if any(
        text_lower == farewell
        or text_lower.startswith(farewell + " ")
        for farewell in farewells
    ):
        return Intent(
            name="farewell",
            confidence=0.95
        )

    question_words = [
        "what",
        "why",
        "how",
        "when",
        "where",
        "who",
        "which",
        "can",
        "could",
        "would",
        "is",
        "are",
        "do",
        "does"
    ]