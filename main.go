package main

import (
	"porter-management/config"
	controller "porter-management/internal/job/application"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDb()

	route := gin.Default()
	controller.InitJobController(route)

	route.Run(":8080")
}
