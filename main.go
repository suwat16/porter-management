package main

import (
	"porter-management/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDb()

	route := gin.Default()

	route.Run(":8080")
}
