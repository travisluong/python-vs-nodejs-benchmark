from fastapi import FastAPI
import asyncpg

app = FastAPI()

@app.on_event("startup")
async def startup():
    global pool
    pool = await asyncpg.create_pool(user='benchmark_user',
                                database='benchmark_db',
                                host='127.0.0.1')

@app.get("/")
async def read_root():
    global pool
    async with pool.acquire() as conn:
        values = await conn.fetch(
            'SELECT * FROM "Post" LIMIT 100'
        )
        return values