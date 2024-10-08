package main

import (
	"github.com/gin-gonic/gin"

	_ "github.com/pienaahj/swag_test/docs"
	"github.com/pienaahj/swag_test/handlers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.

//@securityDefinitions.apikey bearerToken
// @in header
// @name Authorization

// @contact.name API Hendrik Pienaar
// @contact.url https://github.com/pienaahj
// @contact.email pienaahj@gmail.com

// @hostt localhost:8080
// @BasePath /api/v1
func main() {
	// get a router
	r := gin.Default()

	// add user routes
	v1 := r.Group("/api/v1")
	user := v1.Group("/users")
	{
		user.GET("/", handlers.GetUsers)
		// user.GET("/users/:id", GetUser)
		user.POST("/", handlers.CreateUser)
	}

	// add swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// run the router
	err := r.Run()
	if err != nil {
		panic(err)
	}

}
