package api

import (
	"bookstore-api-go/docs"
	authAdmin "bookstore-api-go/pkg/api/admin/auth"
	"bookstore-api-go/pkg/api/books"
	authUser "bookstore-api-go/pkg/api/user/auth"
	profileUser "bookstore-api-go/pkg/api/user/profile"
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
		v1.GET("/", middleware.JWTAuthAdmin(), books.Healthcheck)
		v1.GET("/books", books.FindBooks)
		v1.POST("/books", middleware.JWTAuthUser(), books.CreateBook)
		v1.GET("/books/:id", books.FindBook)
		v1.PUT("/books/:id", books.UpdateBook)
		v1.DELETE("/books/:id", books.DeleteBook)

		v1.POST("/user/login", authUser.LoginHandler)
		v1.POST("/user/register", authUser.RegisterHandler)
		v1.GET("/user/profile", middleware.JWTAuthUser(), profileUser.GetProfile)
		v1.PUT("/user/profile", middleware.JWTAuthUser(), profileUser.UpdateProfile)

		v1.POST("/admin/login", authAdmin.LoginHandler)
		v1.POST("/admin/register", authAdmin.RegisterHandler)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
