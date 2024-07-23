package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json: "id"`
	Title    string `json: "title"`
	Aurthor  string `json: "author"`
	Quantity int    `json: "quantity"`
}

var books = []book{
	{ID: "1", Title: "test1", Aurthor: "test1 Auth", Quantity: 3},
	{ID: "2", Title: "test2", Aurthor: "test2 Auth", Quantity: 5},
	{ID: "3", Title: "test3", Aurthor: "test3 Auth", Quantity: 6},
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.Run("localhost:8080")
}

// curl localhost:8080/books
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// curl localhost:8080/books --include --header "Content-Type: application/json" -d @body.json --request "POST"
func createBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}
