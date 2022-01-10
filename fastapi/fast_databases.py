from typing import List

import databases
from fastapi import FastAPI

database = databases.Database("postgresql://benchmark_user:@localhost/benchmark_db")

app = FastAPI()

@app.on_event("startup")
async def startup():
    await database.connect()

@app.on_event("shutdown")
async def shutdown():
    await database.disconnect()

@app.get("/")
async def root_read():
    query = 'SELECT * FROM "Post" LIMIT 100'
    rows = await database.fetch_all(query=query)
    return rows