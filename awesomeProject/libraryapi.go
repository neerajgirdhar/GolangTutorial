package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID             string `json:"id"`
	Title          string `json:"title"`
	Author         string `json:"author"`
	Quantity       int    `json:"quantity"`
	ActionThriller bool   `json:"actionThriller"`
}

var books = []book{
	{"1", "First Book", "First Author", 5, false},
	{"2", "Second Book", "Second Author", 5, false},
	{"3", "Third Book", "Third Author", 5, false},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func addBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func bookById(c *gin.Context) {

	id := c.Param("id")
	fmt.Printf("Type of ID is %T\n", id)

	book, err := getBookByID(id)
	fmt.Printf("Type of Book ID is %T\n", book.ID)
	fmt.Printf("Type of Book Title is %T\n", book.Title)
	fmt.Printf("Type of Book Quantity is %T\n", book.Quantity)
	fmt.Printf("Type of Book Author is %T\n", book.Author)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func getBookByID(id string) (*book, error) {
	for i, book := range books {
		if book.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("Book Not Found")
}

func bookCheckedOut(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book Id Not passed"})
		return
	}

	book, err := getBookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Found"})
		return
	}
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Available"})
		return
	}

	if book != nil {
		book.Quantity -= 1
	}
	c.IndentedJSON(http.StatusOK, book)
}

func bookCheckedIn(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book Id Not passed"})
		return
	}

	book, err := getBookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Found"})
		return
	}
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Available"})
		return
	}

	if book != nil {
		book.Quantity += 1
	}
	c.IndentedJSON(http.StatusOK, book)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", addBook)
	router.GET("/books/:id", bookById)
	router.PATCH("/bookIssued", bookCheckedOut)
	router.PATCH("/bookReturned", bookCheckedIn)
	router.Run("localhost:8080")

}
