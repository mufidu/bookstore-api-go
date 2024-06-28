package user

import (
	"bookstore-api-go/pkg/database"
	"bookstore-api-go/pkg/models"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

// @BasePath /api/v1

// AddToCart godoc
// @Summary Add a book to the cart
// @Schemes http
// @Description Add a book to the cart of the currently logged in user
// @Tags user
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param   id     body    int     false        "Book ID"
// @Param   quantity     body    int     false        "Quantity"
// @Success 200 {object} models.Cart	"Successfully added book to cart"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Book not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/cart [post]
func AddToCart(c *gin.Context) {
	var cart models.Cart
	var book models.Book
	var cartBook models.CartBook

	// Middleware already sets the username in the context
	username := c.GetString("username")

	// Bind the request body to the cartbook
	if err := c.ShouldBindJSON(&cartBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch the user
	user := models.User{}
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Fetch the cart of the user, if it doesn't exist, create it
	if err := database.DB.Where("user_id = ?", user.ID).Preload("Books").Preload("User").First(&cart).Error; err != nil {
		cart.User = user
		cart.UserID = user.ID
		database.DB.Create(&cart)
	}

	// Add the cart id to the cartbook
	cartBook.CartID = cart.ID

	// Fetch the book
	if err := database.DB.Where("id = ?", cartBook.BookID).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Check if the book is already in the cart
	var existingCartBook models.CartBook
	if err := database.DB.Where("book_id = ? AND cart_id = ?", cartBook.BookID, cart.ID).First(&existingCartBook).Error; err == nil {
		// If the book is already in the cart, update the quantity
		existingCartBook.Quantity = cartBook.Quantity
		cartBook.Quantity = existingCartBook.Quantity
		existingCartBook.BookID = cartBook.BookID
		existingCartBook.CartID = cart.ID
		database.DB.Where("book_id = ? AND cart_id = ?", cartBook.BookID, cart.ID).Save(&existingCartBook)
	} else {
		// If the book is not in the cart, add it
		cart.Books = append(cart.Books, &book)
		cartBook.CartID = cart.ID
		cartBook.BookID = book.ID
		database.DB.Create(&cartBook)
	}

	// Calculate the total price
	cart.TotalPrice = 0
	cartBooks := []models.CartBook{}
	database.DB.Where("cart_id = ?", cart.ID).Find(&cartBooks)
	fmt.Println("CartBooks: ", cartBooks)
	for _, cartBook := range cartBooks {
		book := models.Book{}
		database.DB.Where("id = ?", cartBook.BookID).First(&book)
		cart.TotalPrice += book.Price * cartBook.Quantity
	}

	// Save the cart
	database.DB.Where("id = ?", cart.ID).Save(&cart)

	// Send the updated cart and the cartbook
	c.JSON(http.StatusOK, gin.H{"cart": cart, "cartbook": cartBook})
}

// RemoveFromCart godoc
// @Summary Remove a book from the cart
// @Schemes http
// @Description Remove a book from the cart of the currently logged in user
// @Tags user
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param   id     path    int     true        "Book ID"
// @Success 200 {object} models.Cart	"Successfully removed book from cart"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Book not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/cart/{id} [delete]
func RemoveFromCart(c *gin.Context) {
	var cart models.Cart
	var book models.Book
	var cartBook models.CartBook

	// Middleware already sets the username in the context
	username := c.GetString("username")

	// Get the user id
	user := models.User{}
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Fetch the cart of the user
	if err := database.DB.Where("user_id = ?", user.ID).Preload("Books").Preload("User").First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}

	// Fetch the book
	if err := database.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Fetch the cartbook
	if err := database.DB.Where("book_id = ? AND cart_id = ?", book.ID, cart.ID).First(&cartBook).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not in cart"})
		return
	}

	// Remove the book from the cart
	// database.DB.Where("book_id = ? AND cart_id = ?", book.ID, cart.ID).Delete(&models.CartBook{})

	database.DB.Model(&cart).Association("Books").Delete(&book)

	// If cartbook still exists, throw an error
	if database.DB.Where("book_id = ? AND cart_id = ?", book.ID, cart.ID).First(&cartBook).Error == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove book from cart"})
		return
	}

	// Calculate the total price
	cart.TotalPrice = 0
	cartBooks := []models.CartBook{}
	database.DB.Where("cart_id = ?", cart.ID).Find(&cartBooks)
	for _, cartBook := range cartBooks {
		book := models.Book{}
		database.DB.Where("id = ?", cartBook.BookID).First(&book)
		cart.TotalPrice += book.Price * cartBook.Quantity
	}
	database.DB.Save(&cart)

	// Send the updated cart
	c.JSON(http.StatusOK, gin.H{"Book removed from cart": book})
}

// GetCart godoc
// @Summary Get the cart
// @Schemes http
// @Description Get the cart of the currently logged in user
// @Tags user
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Success 200 {object} models.Cart	"Successfully fetched cart"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Cart not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/cart [get]
func GetCart(c *gin.Context) {
	var cart models.Cart

	// Middleware already sets the username in the context
	username := c.GetString("username")

	// Fetch the user
	user := models.User{}
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Fetch the cart of the user, if it doesn't exist, create it
	if err := database.DB.Where("user_id = ?", user.ID).Preload("Books").Preload("User").First(&cart).Error; err != nil {
		cart.User = user
		cart.UserID = user.ID
		database.DB.Create(&cart)
	}

	// Fetch all cart books
	cartBooks := []models.CartBook{}
	database.DB.Where("cart_id = ?", cart.ID).Find(&cartBooks)

	// Send the cart
	c.JSON(http.StatusOK, gin.H{"cart": cart, "cartbooks": cartBooks})
}

// Checkout godoc
// @Summary Checkout
// @Schemes http
// @Description Checkout the cart of the currently logged in user, and pay for the books with Midtrans (qris as payment method)
// @Tags user
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Success 200 {string} string	"Successfully checked out"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Cart not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/cart/checkout [post]
func Checkout(c *gin.Context) {
	var cart models.Cart
	var cc = coreapi.Client{}
	cc.New(os.Getenv("MIDTRANS_SANDBOX_SERVER_KEY"), midtrans.Sandbox)

	// Middleware already sets the username in the context
	username := c.GetString("username")

	// Fetch the user
	user := models.User{}
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Fetch the cart of the user
	if err := database.DB.Where("user_id = ?", user.ID).Preload("Books").Preload("User").First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}

	// Fetch all cart books
	cartBooks := []models.CartBook{}
	if err := database.DB.Where("cart_id = ?", cart.ID).Find(&cartBooks).Error; err != nil {
		// It means that the cart is empty
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart is empty"})
		return
	}

	// Get the gross amount
	grossAmount := cart.TotalPrice

	// Calculate the service fee
	serviceFee := (grossAmount * 7) / 100

	// Calculate the net amount
	netAmount := grossAmount + serviceFee

	// Get the last transaction
	lastTransaction := models.Transaction{}
	if err := database.DB.Order("id desc").First(&lastTransaction).Error; err != nil {
		lastTransaction.ID = 0
	}

	// Generate the invoice number
	invoiceNumber := "BOOKS-ORDER-0001"
	if lastTransaction.ID != 0 {
		invoiceNumber = "BOOKS-ORDER-" + fmt.Sprintf("%04d", lastTransaction.ID+1)
	}

	// Create string for cart.Books
	items := ""
	for _, cartBook := range cartBooks {
		items += fmt.Sprintf("%d (%d), ", cartBook.BookID, cartBook.Quantity)
	}

	// Create a new transaction
	transaction := models.Transaction{
		UserID:        user.ID,
		InvoiceNumber: invoiceNumber,
		Amount:        netAmount,
		Status:        "pending",
		Items:         items,
	}

	// Create the transaction
	if err := database.DB.Create(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}
	transaction.User = user

	// Create a new charge
	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeQris,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transaction.InvoiceNumber,
			GrossAmt: int64(netAmount),
		},
		CustomerDetails: &midtrans.CustomerDetails{
			Email: user.Email,
		},
	}

	// Charge the transaction
	chargeResp, err := cc.ChargeTransaction(chargeReq)
	if err != nil {
		transaction.Status = "failure"
		database.DB.Save(&transaction)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to charge transaction"})
		return
	}

	// Update the transaction
	transaction.QrisString = chargeResp.QRString
	transaction.ExpiryTime = chargeResp.ExpiryTime
	transaction.InvoiceDate = chargeResp.TransactionTime
	transaction.QrisURL = chargeResp.Actions[0].URL
	transaction.Status = "pending"
	database.DB.Save(&transaction)

	// Remove all cart books that belong to the cart
	database.DB.Model(&cart).Association("Books").Delete(&cartBooks)

	// Empty the cart
	cart.TotalPrice = 0
	database.DB.Save(&cart)

	// Send the charge response
	c.JSON(http.StatusOK, gin.H{"transaction": transaction})
}
