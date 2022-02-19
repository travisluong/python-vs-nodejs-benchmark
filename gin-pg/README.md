# Gin

## Setup

    $ curl https://raw.githubusercontent.com/gin-gonic/examples/master/basic/main.go > main.go
    $ go get -u github.com/gin-gonic/gin
    $ go get -u github.com/lib/pq

## Run benchmark

    $ go run main.go
    $ wrk http://localhost:8080/loadtest