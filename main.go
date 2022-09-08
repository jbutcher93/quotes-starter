package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

type quote struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

var quotes = []quote{
	{Quote: "Don't communicate by sharing memory, share memory by communicating.", Author: "Rob Pike"},
	{Quote: "Concurrency is not parallelism.", Author: "Rob Pike"},
	{Quote: "Clear is better than clever.", Author: "Rob Pike"},
	{Quote: "When reviewing Go code, if I run into a situation where I see an unnecessary deviation from idiomatic Go style or best practice, I add an entry here complete with some rationale, and link to it.", Author: "Dmitri Shuralyov"},
	{Quote: "I can do this for the smallest and most subtle of details, since I care about Go a lot. I can reuse this each time the same issue comes up, instead of having to re-write the rationale multiple times, or skip explaining why I make a given suggestion.", Author: "Dmitri Shuralyov"},
}

func getQuote(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, quotes[rand.Intn(len(quotes)-0)+0])
}

func main() {
	rand.Seed(time.Now().UnixNano())
	router := gin.Default()
	router.GET("/quotes", getQuote)
	router.Run("localhost:8080")
}
