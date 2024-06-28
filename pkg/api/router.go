package api

import (
	"bookstore-api-go/docs"
	"bookstore-api-go/pkg/api/admin"
	"bookstore-api-go/pkg/api/user"
	"bookstore-api-go/pkg/middleware"
	"time"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/time/rate"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())
	if gin.Mode() == gin.ReleaseMode {
		r.Use(middleware.Security())
		r.Use(middleware.Xss())
	}
	r.Use(middleware.Cors())
	r.Use(middleware.RateLimiter(rate.Every(1*time.Minute), 60)) // 60 requests per minute

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		v1.GET("/", admin.Healthcheck)

		v1.POST("/user/login", user.LoginHandler)
		v1.POST("/user/register", user.RegisterHandler)
		v1.GET("/user/profile", middleware.JWTAuthUser(), user.GetProfile)
		v1.PUT("/user/profile", middleware.JWTAuthUser(), user.UpdateProfile)

		v1.GET("/user/books", middleware.JWTAuthUser(), user.GetAllBooks)

		v1.POST("/user/cart", middleware.JWTAuthUser(), user.AddToCart)

		v1.POST("/admin/login", admin.LoginHandler)
		v1.POST("/admin/register", admin.RegisterHandler)
		v1.GET("/admin/users", middleware.JWTAuthAdmin(), admin.GetAllUsers)
		v1.PUT("/admin/users/:username", middleware.JWTAuthAdmin(), admin.UpdateUserByUsername)

		v1.GET("/admin/books", middleware.JWTAuthAdmin(), admin.GetAllBooks)
		v1.POST("/admin/books", middleware.JWTAuthAdmin(), admin.CreateBook)
		v1.PUT("/admin/books/:id", middleware.JWTAuthAdmin(), admin.UpdateBook)
		v1.DELETE("/admin/books/:id", middleware.JWTAuthAdmin(), admin.DeleteBook)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
