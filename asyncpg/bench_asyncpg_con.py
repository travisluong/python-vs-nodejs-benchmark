# https://stackoverflow.com/questions/42242093/asyncpg-connection-vs-connection-pool
import asyncio
import time
import asyncpg

async def bench_asyncpg_con():
    start = time.monotonic()
    for i in range(1, 1000):
        con = await asyncpg.connect(user='benchmark_user', database='benchmark_db',  host='127.0.0.1')
        await con.fetchval('SELECT * FROM "Post" LIMIT 100')
        await con.close()

    end = time.monotonic()
    print(end - start)


asyncio.run(bench_asyncpg_con())