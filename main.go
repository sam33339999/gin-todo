package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"192.168.1.2"})

	// with closure
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// with handler
	router.GET("/test", test)

	// with third-party context passing
	router.GET("/post/:id", GetPost)

	// listen and serve on ... default 8000
	router.Run(":8888")
}

func test(c *gin.Context) {
	str := []byte("hello world")

	c.Data(http.StatusOK, "text/plain", str)
}

func GetPost(ctx *gin.Context) {
	id := ctx.Param("id")

	requestUrl := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%v", id)

	resp, err := http.Get(requestUrl)

	if err != nil {
		log.Println("http.Get failed", err)
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("http.Get failed", err)
		panic(err)
	}

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, string(body))
}
