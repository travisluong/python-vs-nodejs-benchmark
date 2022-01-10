const { Pool, Client } = require('pg')
const express = require('express')
const app = express()
const port = 3000

const client = new Client({
  user: 'benchmark_user',
  host: 'localhost',
  database: 'benchmark_db',
  port: 5432,
})
client.connect()

app.get('/', (req, res) => {
  client.query('SELECT * FROM \"Post\" LIMIT 100', (err, resp) => {
    // console.log(err, resp)
    res.send(resp.rows)
    // client.end()
  })  
})

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
})