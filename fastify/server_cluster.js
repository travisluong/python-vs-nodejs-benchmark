// Require the framework and instantiate it
const fastify = require('fastify')({ logger: false })

const { Pool, Client } = require('pg')
const port = 3000
const cluster = require('cluster');

if (cluster.isMaster) {
    // Count the machine's CPUs
    var cpuCount = require('os').cpus().length;
    console.log(cpuCount)
    // Create a worker for each CPU
    for (var i = 0; i < cpuCount; i += 1) {
        cluster.fork();
    }
} else {
    const client = new Client({
        user: 'benchmark_user',
        host: 'localhost',
        database: 'benchmark_db',
        port: 5432,
      })
      client.connect()
      
      
      // Declare a route
      fastify.get('/', async (request, reply) => {
          const resp = await client.query(`SELECT * FROM "Post" LIMIT 100`);
          return resp.rows;
      //   return { hello: 'world' }
      })
      
      // Run the server!
      const start = async () => {
        try {
          await fastify.listen(3000)
        } catch (err) {
          fastify.log.error(err)
          process.exit(1)
        }
      }
      start()
}

