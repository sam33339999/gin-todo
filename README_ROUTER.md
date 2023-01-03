# Learn Golang - Gin

- router
1. with closure
2. with handler

    - router.setupRouter() *gin..Engine {}

```go
func setupRouter() *gin.Engine {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })
    return r
}

func main() {
    r := setupRouter()
    r.Run() // default localhost:8000
}
```

3. router.parameter -> request -> response

    - http://localhost:8000/post/1
    - http://localhost:8000/post/2 ... etc ...

```go
// router define
router.GET("/post/:id", GetPost)

// handler is here
func GetPost(c *gin.Context) {
    id := c.Param("id")
    requestUrl := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", id)

    resp, err := http.Get(requestURL)
    if err != nil {
        log.Println("http.Get failed", err)
        panic(err)
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    ctx.Header("Content-Type", "application/json")
    ctx.String(http.StatusOK, string(body))
}


```