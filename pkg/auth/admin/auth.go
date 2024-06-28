package auth

import (
	"bookstore-api-go/pkg/database"
	"bookstore-api-go/pkg/models"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Claims struct to be encoded to JWT
type CustomClaims struct {
	Username string `json:"username"`
	UserType string `json:"userType"`
	jwt.StandardClaims
}

var JwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// @BasePath /api/v1

// LoginHandler godoc
// @Summary Authenticate an admin
// @Schemes
// @Description Authenticates an admin using username and password, returns a JWT token if successful
// @Tags admin
// @Accept  json
// @Produce  json
// @Param admin body models.LoginAdmin true "Admin login object"
// @Success 200 {string} string "JWT Token"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/login [post]
func LoginHandler(c *gin.Context) {
	var incomingAdmin models.Admin
	var dbAdmin models.Admin

	// Get JSON body
	if err := c.ShouldBindJSON(&incomingAdmin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	// Fetch the admin from the database
	if err := database.DB.Where("username = ?", incomingAdmin.Username).First(&dbAdmin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(dbAdmin.Password), []byte(incomingAdmin.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate JWT token
	token, err := GenerateToken(dbAdmin.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	// Send admin info and token
	c.JSON(http.StatusOK, gin.H{"admin": dbAdmin, "token": token})
}

// RegisterHandler godoc
// @Summary Register a new admin
// @Schemes http
// @Description Registers a new admin with the given username and password
// @Tags admin
// @Accept  json
// @Produce  json
// @Param   admin     body    models.RegisterAdmin     true        "Admin registration object"
// @Success 200 {string} string	"Successfully registered"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/register [post]
func RegisterHandler(c *gin.Context) {
	var admin models.RegisterAdmin
	var dbAdmin models.Admin

	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password
	hashedPassword, err := HashPassword(admin.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}

	// Create new admin
	newAdmin := models.Admin{Username: admin.Username, Password: hashedPassword, Email: admin.Email}

	// Save the user to the database
	if err := database.DB.Create(&newAdmin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save user"})
		return
	}

	// Generate JWT token
	token, err := GenerateToken(dbAdmin.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	// c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
	c.JSON(http.StatusOK, gin.H{"user": newAdmin, "token": token})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func GenerateToken(username string) (string, error) {
	// The expiration time after which the token will be invalid.
	expirationTime := time.Now().Add(120 * time.Minute).Unix()

	// Create the JWT claims, which includes the username and expiration time
	claims := &CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
			Issuer:    username,
		},
		UserType: "admin",
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(JwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateRandomKey() string {
	key := make([]byte, 32) // generate a 256 bit key
	_, err := rand.Read(key)
	if err != nil {
		panic("Failed to generate random key: " + err.Error())
	}

	return base64.StdEncoding.EncodeToString(key)
}
