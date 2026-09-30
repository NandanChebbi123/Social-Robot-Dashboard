from pydantic import BaseModel, Field
from typing import Optional
from datetime import datetime

class Message(BaseModel):
    speaker: str
    content: str
    timestamp: str

class ConversationSession(BaseModel):
    session_id: str

    messages: list[Message] = Field(
        default_factory = list
    )

    created_at: str
    last_updated: str

class ContextManager:

    def __init__(self):

        self.sessions: dict[str, ConversationSession] = {}

    def create_session(
            self,
            session_id: str
    ):
        
        now = datetime.now().isoformat()

        session = ConversationSession(
            session_id=session_id,
            created_at=now,
            last_updated=now
        )

        self.sessions[session_id] = session

        return session

    def get_session(
            self,
            session_id:str
    ):

        if session_id not in self.sessions:

            return self.create_session(
                session_id
            )

        return self.sessions[session_id]

    def add_user_message(
            self,
            session_id: str,
            content: str
    ):

        session = self.get_session(
            session_id
        )

        message = Message(
            speaker="user",
            content=content,
            timestamp=datetime.now().isoformat()
        )

        session.messages.append(message)

        session.last_updated = (
            datetime.now().isoformat()
        )

    def add_robot_message(
            self,
            session_id: str,
            content: str
    ):

        session = self.get_session(
            session_id
        )

        message = Message(
            speaker="robot",
            content=content,
            timestamp=datetime.now().isoformat()
        )

        session.messages.append(message)

        session.last_updated = (
            datetime.now().isoformat()
        )

        def get_recent_messages(
                self,
                session_id:str,
                limit: int = 10
        ):

            session = self.get_session(
                session_id
            )

            return session.messages[-limit:]

        def build_context(
                self,
                session_id: str,
                current_input: str,
                limit: int = 10
        ):

            recent_messages = self.get_recent_messages(
                session_id,
                limit
            )

            return Context(
                session_id = session_id,
                current_input=current_input,
                recent_messages=recent_messages
            )

        def clear_sessions(
                self,
                session_id: str
        ):

            if session_id in self.sessions:
                
                del self.sessions[session_id]