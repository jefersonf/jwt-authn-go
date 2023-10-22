package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jefersonf/jwt-authn-go/env"
	"github.com/jefersonf/jwt-authn-go/middleware"
	"github.com/jefersonf/jwt-authn-go/user"
)

var (
	listenAddr string
)

func init() {
	env.LoadEnvVars()

	// db connector
	conn := env.Conn()

	env.SyncDatabase(conn())
}

func main() {
	routes := gin.Default()
	routes.POST("/signup", user.Signup)
	routes.POST("/login", user.Login)
	routes.GET("/validate", middleware.EnsureAuth, user.Validate)
	fmt.Println("server API running..")
	routes.Run()
}
