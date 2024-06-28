package main

import (
	"bookstore-api-go/pkg/api"
	"bookstore-api-go/pkg/cache"
	"bookstore-api-go/pkg/database"
	"log"

	"github.com/gin-gonic/gin"
)

// @title           Bookstore API GO
// @version         1.0
// @description     API for bookstore.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:9000
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	cache.InitRedis()
	database.ConnectDatabase()

	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)

	r := api.InitRouter()

	if err := r.Run(":9000"); err != nil {
		log.Fatal(err)
	}
}
