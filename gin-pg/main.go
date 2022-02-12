package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	// "encoding/json"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "benchmark_user"
	// password = ""
	dbname = "benchmark_db"
)

type Post struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Published bool   `json:"published"`
}

type JsonResponse struct {
	Data []Post `json:"data"`
}

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)

	pg, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	pg.SetMaxOpenConns(10000)
	pg.SetMaxIdleConns(5000)

	r.GET("/hello", func(c *gin.Context) {
		rows, err := pg.Query("SELECT id, content FROM \"Post\" LIMIT 100;")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		log.Print(rows)
		for rows.Next() {
			var Id int
			var Content string
			// var published bool
			err = rows.Scan(&Id, &Content)
			if err != nil {
				panic(err)
			}
			// fmt.Println(Id, Content)
		}

		c.String(http.StatusOK, "hello")
	})

	r.GET("/loadtest", func(c *gin.Context) {
		rows, err := pg.Query("SELECT id, title, content, published FROM \"Post\" LIMIT 100;")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		// log.Print(rows)
		posts := make([]Post, 0)
		for rows.Next() {
			var Id int
			var Title string
			var Content string
			var Published bool
			err = rows.Scan(
				&Id,
				&Title,
				&Content,
				&Published,
			)
			if err != nil {
				panic(err)
			}
			// fmt.Println(Id, Title, Content, Published)
			posts = append(posts, Post{Id: Id, Title: Title, Content: Content, Published: Published})
			// fmt.Println(Id, Content)
		}
		// log.Print(posts)
		// var response = JsonResponse{Data: posts}
		// log.Print(response)
		// out, err := json.Marshal(posts)
		// if err != nil {
		// panic(err)
		// }
		// log.Print(out)
		// c.String(http.StatusOK, "hello")
		c.JSON(http.StatusOK, gin.H{"posts": posts})
	})

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	log.Print("HELLO WORLD!!!!")

	r := setupRouter()

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
