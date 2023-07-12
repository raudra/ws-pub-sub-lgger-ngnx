package main

import (
	"net/http"
	middleware "profile-service/middlewares"
	"profile-service/src/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Use(middleware.DefaultStructuredLogger()) // adds our new middleware
	router.Use(middleware.LogResponseBody)
	router.Use(gin.Recovery())
	router.GET("ping", getRoute1)

	v1 := router.Group("/api/v1")
	{
		v1.POST("/users", controllers.CreateUser)
		v1.GET("/users/:id", controllers.GetUser)
		v1.GET("/users/findByMobile", controllers.GetUserByMobile)
	}

	router.Run()
}

func getRoute1(c *gin.Context) {
	data := map[string]string{
		"status": "ok",
	}
	c.JSON(http.StatusOK, data)
}
