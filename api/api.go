package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
)

// book represents data about a record book.
type book struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Author string  `json:"author"`
    Cost  float64 `json:"price"`
    Quantity int `json:"qty"`
}

// books slice to seed record book data.
var books = []book {
    {ID: "1", Title: "Blue Train", Author: "John Coltrane", Cost: 56.99, Quantity: 2},
    {ID: "2", Title: "Jeru", Author: "Gerry Mulligan", Cost: 27.99, Quantity: 1},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Author: "Sarah Vaughan", Cost: 39.99, Quantity: 1},
}

// getAlbums responds with the list of all books as JSON.
func getBooks(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, books)
}

// returns a pointer and error
func getBookById(id string) (*book, error) {
    for i, a := range books {
        if a.ID == id {
            return &books[i], nil
        }
    }
    return nil, errors.New("book now found")
}

func bookById(c *gin.Context) {
    id := c.Param("id")
    book, err := getBookById(id)

    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
        return
    }

    c.IndentedJSON(http.StatusOK, book)
}

// createBook responds with a created book
func createBook(c *gin.Context) {
    var newBook book

    if err := c.BindJSON(&newBook); err != nil {
        return
    }

    books = append(books, newBook)
    c.IndentedJSON(http.StatusCreated, newBook)
}

func checkoutBook(c * gin.Context) {
    id, ok := c.GetQuery("id")

    if ok == false {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id query parameter"})
        return
    }

    book, err := getBookById(id)

    if err != nil {
        response := "The book" + book.Title + "was not found."
        c.IndentedJSON(http.StatusNotFound, gin.H{ "message": response })
        return
    }

   if book.Quantity <= 0 {
    c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available"})
    return
   }

   book.Quantity -= 1
   c.IndentedJSON(http.StatusOK, book)
}

func main() {
    router := gin.Default()
    router.GET("/books", getBooks)
    router.GET("/book/:id", bookById)
    router.POST("/book", createBook)
    router.PATCH("/checkout", checkoutBook)
    router.Run("localhost:8080")
}