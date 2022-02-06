from flask import Flask
import psycopg2
import psycopg2.extras
import orjson

app = Flask(__name__)
conn = psycopg2.connect("dbname=benchmark_db user=benchmark_user")
cur = conn.cursor(cursor_factory=psycopg2.extras.RealDictCursor)

@app.route("/")
def hello_world():
    cur.execute("SELECT * FROM \"Post\" LIMIT 100")
    records = cur.fetchall()
    return orjson.dumps(records)