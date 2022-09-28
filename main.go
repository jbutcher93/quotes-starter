package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type quote struct {
	ID     string `json:"id"`
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

type id struct {
	ID string `json:"id"`
}

var quotes = map[string]quote{
	"b513f4ec-ddd8-4d54-ae47-d78a2ab43612": {ID: "b513f4ec-ddd8-4d54-ae47-d78a2ab43612", Quote: "Don't communicate by sharing memory, share memory by communicating.", Author: "Rob Pike"},
	"dd50bc8d-17cc-4bf9-a5ce-674d9e501408": {ID: "dd50bc8d-17cc-4bf9-a5ce-674d9e501408", Quote: "Concurrency is not parallelism.", Author: "Rob Pike"},
	"fe4a522e-d668-4e51-85aa-65be75049618": {ID: "fe4a522e-d668-4e51-85aa-65be75049618", Quote: "Clear is better than clever.", Author: "Rob Pike"},
	"20f41f99-159d-4539-9182-3c3e531ebbf9": {ID: "20f41f99-159d-4539-9182-3c3e531ebbf9", Quote: "When reviewing Go code, if I run into a situation where I see an unnecessary deviation from idiomatic Go style or best practice, I add an entry here complete with some rationale, and link to it.", Author: "Dmitri Shuralyov"},
	"b0372040-94aa-4f1f-b905-05ee2e2efecd": {ID: "b0372040-94aa-4f1f-b905-05ee2e2efecd", Quote: "I can do this for the smallest and most subtle of details, since I care about Go a lot. I can reuse this each time the same issue comes up, instead of having to re-write the rationale multiple times, or skip explaining why I make a given suggestion.", Author: "Dmitri Shuralyov"},
}

func getRandomQuote() *quote {
	randNum := (rand.Intn(len(quotes)))
	counter := 0

	for _, v := range quotes {
		if randNum == counter {
			return &v
		}
		counter++
	}
	return nil
}

func returnRandomQuote(c *gin.Context) {
	if validateKey(c) != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, getRandomQuote())
}

func getQuoteWithID(id string) (quote, error) {
	v, ok := quotes[id]
	if !ok {
		return quote{}, errors.New("this quote ID does not match any of our records")
	}
	return v, nil
}

func returnQuoteWithId(c *gin.Context) {
	if validateKey(c) != nil {
		return
	}

	id := c.Param("id")
	quote, err := getQuoteWithID(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		return
	}

	c.IndentedJSON(http.StatusOK, quote)
}

func postQuote(c *gin.Context) {
	if validateKey(c) != nil {
		return
	}

	var newQuote quote
	err := c.BindJSON(&newQuote)
	validator := validateStruct(&newQuote)

	if err != nil {
		fmt.Println(errors.New("failed to create new quote"))
		c.AbortWithStatus(400)
		return
	}

	if len(validator) > 0 {
		for _, v := range validator {
			fmt.Println(v)
		}
		c.AbortWithStatus(400)
		return
	}

	newQuote.ID = createID()
	quotes[newQuote.ID] = newQuote
	id := id{newQuote.ID}
	c.IndentedJSON(http.StatusCreated, id)
}

func validateStruct(q *quote) []error {
	err := make([]error, 0, 2)

	if len(q.Author) < 3 {
		err = append(err, errors.New("author must be at least 3 characters"))
	}
	if len(q.Quote) < 3 {
		err = append(err, errors.New("quote must be at least 3 characters"))
	}

	return err
}

func validateKey(c *gin.Context) error {
	token := strings.Join(c.Request.Header["X-Api-Key"], "")

	if token != "COCKTAILSAUCE" {
		c.AbortWithStatus(401)
		return errors.New("not COCKTAILSAUCE")
	}

	return nil
}

func createID() string {
	id := uuid.New().String()
	return id
}

func main() {
	rand.Seed(time.Now().UnixNano())
	router := gin.Default()
	router.GET("/quotes", returnRandomQuote)
	router.GET("/quotes/:id", returnQuoteWithId)
	router.POST("/quotes", postQuote)
	router.Run("0.0.0.0:8080")
}
