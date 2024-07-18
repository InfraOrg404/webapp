package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	//initialize a db connection
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//comment
	//gin router
	router := gin.Default()
	router.GET("/healthz", healthzhHandler)
	router.NoRoute(methodNotAllowedHandler)

	//start the server on port 8080
	router.Run(":8080")
}

func healthzhHandler(c *gin.Context) {
	//check for any payload in the request
	if c.Request.ContentLength > 0 {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	//check for db connectivity
	sqlDB, err := db.DB()
	if err != nil || sqlDB.Ping() != nil {
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Header("Pragma", "no-cache")
		c.Header("X-Content-Type-Options", "nosniff")
		c.JSON(http.StatusOK, gin.H{})
	}
}

func methodNotAllowedHandler(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header("Pragma", "no-cache")
	c.Header("X-Content-Type-Options", "nosniff")
	c.JSON(http.StatusMethodNotAllowed, gin.H{})
}
