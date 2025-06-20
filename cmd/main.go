package main

import (
	"clinic-portal/config"
	"clinic-portal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":" + config.EnvPort())
}
