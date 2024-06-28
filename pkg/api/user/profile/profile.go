package profileUser

import (
	"bookstore-api-go/pkg/database"
	"bookstore-api-go/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProfile godoc
// @Summary Get user profile
// @Schemes http
// @Description Get the profile of the currently logged in user
// @Tags user
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Success 200 {object} models.User	"Successfully retrieved user profile"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/profile [get]
func GetProfile(c *gin.Context) {
	var user models.User

	// Middleware already sets the username in the context
	username := c.GetString("username")

	// Fetch the user from the database
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		// not found
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Return the user's profile
	c.JSON(http.StatusOK, gin.H{"user": user})
}
