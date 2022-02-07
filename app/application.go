package app

import (
	"github.com/ceeghee/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	// MapUrls()
	GroupUrls()
	logger.Info("About to start the application")
	router.Run(":8080")
}
