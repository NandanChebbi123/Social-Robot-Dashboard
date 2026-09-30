from pydantic import BaseModel
from typic import Optional

class RawInput(BaseModel):
    session_id: str
    content: str
    timestap: Optional[str] = None

class ProcessedInput(BaseModel):
    session_id: str

    raw_content: str
    normalized_content: str

    timestamp: Optional[str] = None

def validate_input(data: RawInput):
    if not data.content.strip():
        raise ValueError("Input cannot be empty")

def normalize_text(text:str) -> str:
    return " ".join(text.split())

def process_input(data: RawInput) -> ProcessedInput:

    validate_input(data)

    normalized = normalize_text(data.content)

    return ProcessedInput(
        session_id=data.session_id,
        raw_content=data.content,
        normalized_content=normalized,
        timestamp=data.timestap
    )