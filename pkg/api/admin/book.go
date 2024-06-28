package admin

import (
	"bookstore-api-go/pkg/database"
	"bookstore-api-go/pkg/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// Healthcheck godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router / [get]
func Healthcheck(g *gin.Context) {
	g.JSON(http.StatusOK, "ok")
}

// GetAllBooks godoc
// @Summary Get all books
// @Schemes http
// @Description Get all books
// @Tags admin
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Success 200 {object} models.Book	"Successfully retrieved all books"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/books [get]
func GetAllBooks(c *gin.Context) {
	var books []models.Book

	if err := database.DB.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"books": books})
}

// CreateBook godoc
// @Summary Create a book
// @Schemes http
// @Description Create a book
// @Tags admin
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param   book     body    models.CreateBook     true        "Book object"
// @Success 201 {object} models.Book	"Successfully created book"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/books [post]
func CreateBook(c *gin.Context) {
	var book models.CreateBook

	// Print the incoming JSON body
	fmt.Println(c.Request.Body)

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	newBook := models.Book{
		Title:    book.Title,
		Price:    book.Price,
		Genre:    book.Genre,
		Cover:    book.Cover,
		Author:   book.Author,
		Year:     book.Year,
		Quantity: book.Quantity,
	}

	if err := database.DB.Create(&newBook).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"book": newBook})
}

// UpdateBook godoc
// @Summary Update a book
// @Schemes http
// @Description Update a book
// @Tags admin
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param   id     path    int     true        "Book ID"
// @Param   book     body    models.UpdateBook     true        "Book object"
// @Success 200 {object} models.Book	"Successfully updated book"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/books/{id} [put]
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.UpdateBook

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	var existingBook models.Book
	if err := database.DB.Where("id = ?", id).First(&existingBook).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}

	existingBook.Title = book.Title
	existingBook.Price = book.Price
	existingBook.Genre = book.Genre
	existingBook.Cover = book.Cover
	existingBook.Author = book.Author
	existingBook.Year = book.Year
	existingBook.Quantity = book.Quantity

	if err := database.DB.Save(&existingBook).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"book": existingBook})
}

// DeleteBook godoc
// @Summary Delete a book
// @Schemes http
// @Description Delete a book
// @Tags admin
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param   id     path    int     true        "Book ID"
// @Success 204 {string} string	"Successfully deleted book"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/books/{id} [delete]
func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book
	if err := database.DB.Where("id = ?", id).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}

	if err := database.DB.Delete(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted book"})
}
