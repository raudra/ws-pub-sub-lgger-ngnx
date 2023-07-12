package main

import (
	"net/http"
	middleware "session-service/middlewares"
	"session-service/src/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Use(middleware.DefaultStructuredLogger()) // adds our new middleware
	router.Use(gin.Recovery())
	router.Use(middleware.LogResponseBody)
	router.GET("ping", getRoute1)

	v1 := router.Group("/api/v1")
	{
		v1.POST("/sessions", controllers.CreateSession)
		v1.DELETE("/sessions", controllers.DeleteSession)
	}

	router.Run()
}

func getRoute1(c *gin.Context) {
	data := map[string]string{
		"status": "ok",
	}
	c.JSON(http.StatusOK, data)
}
