package admin

import (
	"bookstore-api-go/pkg/database"
	"bookstore-api-go/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// GetAllUsers godoc
// @Summary Get all users
// @Schemes http
// @Description Get all users
// @Tags admin
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Success 200 {object} models.User	"Successfully retrieved all users"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/users [get]
func GetAllUsers(c *gin.Context) {
	var users []models.User

	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// UpdateUserByUsername godoc
// @Summary Update user by username
// @Schemes http
// @Description Update user by username
// @Tags admin
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param   username     path    string     true        "Username"
// @Param   user     body    models.RegisterUser     true        "User object"
// @Success 200 {object} models.User	"Successfully updated user"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/users/{username} [put]
func UpdateUserByUsername(c *gin.Context) {
	var user models.User
	username := c.Param("username")

	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
