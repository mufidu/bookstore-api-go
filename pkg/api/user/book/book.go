package bookUser

import (
	"bookstore-api-go/pkg/database"
	"bookstore-api-go/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// GetAllBooks godoc
// @Summary Get all books
// @Schemes http
// @Description Get all books
// @Tags user
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param genre query string false "Genre"
// @Param author query string false "Author"
// @Param year query int false "Year"
// @Success 200 {object} models.Book	"Successfully retrieved all books"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/books [get]
func GetAllBooks(c *gin.Context) {
	var books []models.Book

	// Retrieve all books
	if err := database.DB.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Filter books by genre
	if genre := c.Query("genre"); genre != "" {
		if err := database.DB.Where("genre = ?", genre).Find(&books).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
	}

	// Filter books by author
	if author := c.Query("author"); author != "" {
		if err := database.DB.Where("author = ?", author).Find(&books).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
	}

	// Filter books by year
	if year := c.Query("year"); year != "" {
		if err := database.DB.Where("year = ?", year).Find(&books).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"books": books})
}
