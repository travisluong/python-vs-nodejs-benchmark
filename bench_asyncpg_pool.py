# https://stackoverflow.com/questions/42242093/asyncpg-connection-vs-connection-pool
import asyncio
import time
import asyncpg

async def bench_asyncpg_pool():
    pool = await asyncpg.create_pool(user='benchmark_user', database='benchmark_db', host='127.0.0.1')
    start = time.monotonic()
    for i in range(1, 1000):
        async with pool.acquire() as con:
            await con.fetchval('SELECT * FROM "Post" LIMIT 100')

    await pool.close()
    end = time.monotonic()
    print(end - start)

asyncio.run(bench_asyncpg_pool())