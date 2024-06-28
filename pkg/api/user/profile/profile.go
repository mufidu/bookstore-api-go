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

// UpdateProfile godoc
// @Summary Update user profile
// @Schemes http
// @Description Update the profile of the currently logged in user
// @Tags user
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param   user     body    models.RegisterUser     true        "User object"
// @Success 200 {object} models.User	"Successfully updated user profile"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/profile [put]
func UpdateProfile(c *gin.Context) {
	var user models.User

	// Middleware already sets the username in the context
	username := c.GetString("username")

	// Get JSON body
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the user in the database
	if err := database.DB.Model(&user).Where("username = ?", username).Updates(user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Return the updated user's profile
	c.JSON(http.StatusOK, gin.H{"user": user})
}
