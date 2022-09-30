package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v4/stdlib"
	// Note: If connecting using the App Engine Flex Go runtime, use
	// "github.com/jackc/pgx/stdlib" instead, since v4 requires
	// Go modules which are not supported by App Engine Flex.
)

type quote struct {
	ID     string `json:"id"`
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

type id struct {
	ID string `json:"id"`
}

var db *sql.DB

func databaseConnection() error {
	mustGetenv := func(dns string) string {
		gettingEnv := os.Getenv(dns)
		if gettingEnv == "" {
			log.Printf("Warning: %s environment variable not set", dns)
		}
		return gettingEnv
	}

	var (
		dbUser         = os.Getenv("DB_USER") //postgres
		dbPwd          = mustGetenv("DB_PASS")
		dbName         = mustGetenv("DB_NAME") //quotes_database
		unixSocketPath = mustGetenv("INSTANCE_UNIX_SOCKET")
	)

	dbURI := fmt.Sprintf("user=%s password=%s database=%s host=%s", dbUser, dbPwd, dbName, unixSocketPath)

	//dbPool is the pool of database connections
	var err error

	db, err = sql.Open("pgx", dbURI)
	if err != nil {
		return fmt.Errorf("sql.Open: %v", err)
	}
	return err
}

func getRandomQuote() *quote {
	row := db.QueryRow("SELECT * FROM quotes ORDER BY RANDOM () LIMIT 1;")
	q := &quote{}
	err := row.Scan(&q.ID, &q.Quote, &q.Author)
	if err != nil {
		log.Println(err)
	}
	return q
}

func returnRandomQuote(c *gin.Context) {
	if validateKey(c) != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, getRandomQuote())
}

func getQuoteWithId(id string) *quote {
	row := db.QueryRow(fmt.Sprintf("select * from quotes where ID = '%s'", id))
	q := &quote{}
	err := row.Scan(&q.ID, &q.Quote, &q.Author)
	if err != nil {
		log.Println(err)
	}
	return q
}

func returnQuoteWithId(c *gin.Context) {
	if validateKey(c) != nil {
		return
	}

	id := c.Param("id")
	quote := getQuoteWithId(id)

	if quote.ID == "" {
		c.AbortWithStatus(404)
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
	db.Exec("INSERT INTO quotes (id, author, phrase) VALUES ($1, $2, $3);", &newQuote.ID, &newQuote.Author, &newQuote.Quote)
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
	err := databaseConnection()
	if err != nil {
		log.Println(err)
	}
	rand.Seed(time.Now().UnixNano())
	router := gin.Default()
	router.GET("/quotes", returnRandomQuote)
	router.GET("/quotes/:id", returnQuoteWithId)
	router.POST("/quotes", postQuote)
	router.Run("0.0.0.0:8080")
}
