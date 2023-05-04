package main

import (
	"github.com/amal-007-amal/auth-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	routes.UserRoute()
}
