# Flask

## Setup

    $ pip3 install -r requirements.txt

## Run

    $ FLASK_APP=flask_psycopg flask run
    $ wrk http://localhost:5000

    $ gunicorn -w 1 --bind 0.0.0.0:5000 flask_psycopg:app
    $ wrk http://localhost:5000

    $ gunicorn -w 4 --bind 0.0.0.0:5000 flask_psycopg:app
    $ wrk http://localhost:5000