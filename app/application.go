package app

import (
	"github.com/JCFlores93/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("about to start the application...")
	router.Run(":8082")
	logger.Log.Info("about to finish the application...")
}