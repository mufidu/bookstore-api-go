package user

import (
	"bookstore-api-go/pkg/database"
	"bookstore-api-go/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Add to Cart: Customers can add books to their shopping cart.
// Update Cart: Customers can update the quantity or remove items from their cart.
// Total Price Calculation: The total price of items in the cart is automatically calculated each time the cart is updated.

// @BasePath /api/v1

// AddToCart godoc
// @Summary Add a book to the cart
// @Schemes http
// @Description Add a book to the cart of the currently logged in user
// @Tags user
// @Accept  json
// @Produce  json
// @Security BearerAuth
// Should ask for the book id and the quantity in body, not the cartbook
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

	// Fetch the cart of the user, if it doesn't exist, create it
	if err := database.DB.Where("user_id = ?", username).First(&cart).Error; err != nil {
		user := models.User{}
		if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		cart.UserID = user.ID
		database.DB.Create(&cart)
	}

	// Fetch the book
	if err := database.DB.Where("id = ?", cartBook.BookID).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Check if the book is already in the cart
	var existingCartBook models.CartBook
	if err := database.DB.Where("book_id = ? AND cart_id = ?", cartBook.BookID, cart.ID).First(&existingCartBook).Error; err == nil {
		// If the book is already in the cart, update the quantity
		existingCartBook.Quantity += cartBook.Quantity
		database.DB.Save(&existingCartBook)
	} else {
		// If the book is not in the cart, add it
		cart.Books = append(cart.Books, &book)
		cartBook.CartID = cart.ID
		database.DB.Create(&cartBook)
	}

	// Calculate the total price
	cart.TotalPrice += book.Price * cartBook.Quantity
	database.DB.Save(&cart)

	// Send the updated cart and the cartbook
	c.JSON(http.StatusOK, gin.H{"cart": cart, "cartbook": cartBook})
}
