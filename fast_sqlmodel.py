from typing import Optional
from fastapi import FastAPI
from sqlmodel import Field, SQLModel, Session, create_engine, select

app = FastAPI()

class Post(SQLModel, table=True):
    __tablename__ = "Post"

    id: Optional[int] = Field(default=None, primary_key=True)
    title: str
    content: str
    published: bool

engine = create_engine("postgresql://benchmark_user:@localhost/benchmark_db")

@app.get("/")
async def read_root():
    with Session(engine) as session:
        notes = session.exec(select(Post).limit(100)).all()
        return notes