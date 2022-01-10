import psycopg2
import random

conn = psycopg2.connect("dbname=benchmark_db user=benchmark_user")
cur = conn.cursor()

for i in range(100):
    cur.execute("INSERT INTO \"Post\" (title, content, published) VALUES (%s, %s, %s)", ('foobar', 'foobar', True))

conn.commit()