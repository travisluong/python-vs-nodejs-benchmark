# FastAPI

## Setup

    $ pip3 install -r requirements.txt

## Run benchmark

    $ uvicorn fast_psycopg:app
    $ wrk http://localhost:8000

    $ uvicorn fast_sqlmodel:app
    $ wrk http://localhost:8000

    $ gunicorn -w 4 -k uvicorn.workers.UvicornWorker fast_psycopg:app
    $ wrk http://localhost:8000

    $ gunicorn -w 4 -k uvicorn.workers.UvicornWorker fast_sqlmodel:app
    $ wrk http://localhost:8000

    $ uvicorn fast_asyncpg:app
    $ wrk http://localhost:8000

    $ gunicorn -w 4 -k uvicorn.workers.UvicornWorker fast_asyncpg:app
    $ wrk http://localhost:8000

    $ uvicorn fast_databases:app
    $ wrk http://localhost:8000

    $ uvicorn fast_psycopg:app
    $ wrk http://localhost:8000/orjson

    $ uvicorn fast_asyncpg:app
    $ wrk http://localhost:8000/orjson

    $ uvicorn fast_asyncpg:app
    $ wrk http://localhost:8000/ujson

    $ gunicorn -w 4 -k uvicorn.workers.UvicornWorker fast_asyncpg:app
    $ wrk http://localhost:8000/orjson

    $ gunicorn -w 4 -k uvicorn.workers.UvicornWorker fast_asyncpg:app
    $ wrk http://localhost:8000/ujson