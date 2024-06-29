package admin

import (
	"bookstore-api-go/pkg/database"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	// Load .env file from the root directory
	if err := godotenv.Load("../../../.env"); err != nil {
		fmt.Println("No .env file found")
	}

	gin.SetMode(gin.TestMode)
	database.ConnectDatabase()
}

func GetAdminToken() string {
	// Direct JSON string payload
	jsonPayload := `{
        "username": "admin",
        "password": "admin123!",
    }`

	router := gin.Default()

	router.POST("/api/v1/admin/login", LoginHandler)

	// Use strings.NewReader to create the request body from jsonPayload
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/admin/login", strings.NewReader(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	// Print request body and status code
	fmt.Println("HTTP Method:", req.Method)
	fmt.Println("HTTP URI:", req.RequestURI)
	fmt.Println("HTTP Header:", req.Header)
	fmt.Println("HTTP Body:", strings.NewReader(jsonPayload))
	fmt.Println("HTTP Status Code:", w.Code)
	fmt.Println("BODY:", w.Body)

	// Define a struct to match the JSON response format
	type loginResponse struct {
		Admin struct {
			ID        int    `json:"id"`
			Username  string `json:"username"`
			Email     string `json:"email"`
			Password  string `json:"password"`
			CreatedAt string `json:"created_at"`
			UpdatedAt string `json:"updated_at"`
		} `json:"admin"`
		Token string `json:"token"`
	}

	// Parse the JSON response to extract the token
	var resp loginResponse
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		fmt.Println("Error parsing JSON:", err)
	}

	return resp.Token
}

func TestMainRouter(t *testing.T) {
	// Set up Gin
	router := gin.Default()

	// Register the endpoint to test
	router.GET("/api/v1", Healthcheck)

	// Create a test request to the endpoint
	req, _ := http.NewRequest(http.MethodGet, "/api/v1", nil)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetAllUsers(t *testing.T) {
	// Set up Gin
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Register the endpoint to test
	router.GET("/api/v1/admin/users", GetAllUsers)

	// Get admin token
	token := GetAdminToken()

	// Create a test request to the endpoint
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/admin/users", nil)
	// Add the token to the Authorization header
	req.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateUserByUsername(t *testing.T) {
	// Set up Gin
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Register the endpoint to test
	router.PUT("/api/v1/admin/users/:username", UpdateUserByUsername)

	// Get admin token
	token := GetAdminToken()

	// Create a test request to the endpoint
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/admin/users/user", nil)
	// Add the token to the Authorization header
	req.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetAllBooks(t *testing.T) {
	// Set up Gin
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Register the endpoint to test
	router.GET("/api/v1/admin/books", GetAllBooks)

	// Get admin token
	token := GetAdminToken()

	// Create a test request to the endpoint
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/admin/books", nil)
	// Add the token to the Authorization header
	req.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateBook(t *testing.T) {
	// Set up Gin
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Register the endpoint to test
	router.POST("/api/v1/admin/books", CreateBook)

	// Get admin token
	token := GetAdminToken()

	// Create a test request to the endpoint
	jsonPayload := `{
		"title": "Test Book",
		"author": "Test Author",
		"price": 10.99,
		"quantity": 10
	}`
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/admin/books", strings.NewReader(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	// Add the token to the Authorization header
	req.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateBook(t *testing.T) {
	// Set up Gin
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Register the endpoint to test
	router.PUT("/api/v1/admin/books/:id", UpdateBook)

	// Get admin token
	token := GetAdminToken()

	// Create a test request to the endpoint
	jsonPayload := `{
		"title": "Test Book",
		"author": "Test Author",
		"price": 10.99,
		"quantity": 10
	}`
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/admin/books/1", strings.NewReader(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	// Add the token to the Authorization header
	req.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteBook(t *testing.T) {
	// Set up Gin
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Register the endpoint to test
	router.DELETE("/api/v1/admin/books/:id", DeleteBook)

	// Get admin token
	token := GetAdminToken()

	// Create a test request to the endpoint
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/admin/books/1", nil)
	// Add the token to the Authorization header
	req.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
}
