package api

import (
	"bookstore-api-go/docs"
	authAdmin "bookstore-api-go/pkg/api/admin/auth"
	bookAdmin "bookstore-api-go/pkg/api/admin/book"
	userAdmin "bookstore-api-go/pkg/api/admin/user"
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
		v1.GET("/", bookAdmin.Healthcheck)

		v1.POST("/user/login", authUser.LoginHandler)
		v1.POST("/user/register", authUser.RegisterHandler)
		v1.GET("/user/profile", middleware.JWTAuthUser(), profileUser.GetProfile)
		v1.PUT("/user/profile", middleware.JWTAuthUser(), profileUser.UpdateProfile)

		v1.POST("/admin/login", authAdmin.LoginHandler)
		v1.POST("/admin/register", authAdmin.RegisterHandler)
		v1.GET("/admin/users", middleware.JWTAuthAdmin(), userAdmin.GetAllUsers)
		v1.PUT("/admin/users/:username", middleware.JWTAuthAdmin(), userAdmin.UpdateUserByUsername)

		v1.GET("/admin/books", middleware.JWTAuthAdmin(), bookAdmin.GetAllBooks)
		v1.POST("/admin/books", middleware.JWTAuthAdmin(), bookAdmin.CreateBook)
		v1.PUT("/admin/books/:id", middleware.JWTAuthAdmin(), bookAdmin.UpdateBook)
		v1.DELETE("/admin/books/:id", middleware.JWTAuthAdmin(), bookAdmin.DeleteBook)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
