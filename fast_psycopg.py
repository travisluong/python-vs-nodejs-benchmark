from typing import Optional
from fastapi import FastAPI
import psycopg2
import psycopg2.extras

app = FastAPI()
conn = psycopg2.connect("dbname=benchmark_db user=benchmark_user")
cur = conn.cursor(cursor_factory=psycopg2.extras.RealDictCursor)

@app.get("/")
async def read_root():
    cur.execute("SELECT * FROM \"Post\" LIMIT 100")
    records = cur.fetchall()
    return records