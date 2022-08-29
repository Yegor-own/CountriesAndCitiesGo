package main

import (
	"countries-and-cities/app/controllers"
	"github.com/gin-gonic/gin"
)

func api() {
	router := gin.Default()

	router.GET("/api/countries", controllers.CountryIndex)
	router.GET("/api/countries/:id")
	router.POST("/api/countries")
	router.PUT("/api/countries/:id")
	router.DELETE("/api/countries/:id")

	router.GET("/api/cities")
	router.GET("/api/cities/:id")
	router.POST("/api/cities")
	router.PUT("/api/cities/:id")
	router.DELETE("/api/cities/:id")

	router.Run(":3000")
}
