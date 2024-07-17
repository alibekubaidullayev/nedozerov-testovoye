package main

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	port string = "6060"
)

var db []Account = make([]Account, 0)

func main() {
	router := gin.New()
	router.Use(gin.Recovery())

	RegAccountsRouter(router, "accounts")

	if err := router.Run(":" + port); err != nil {
		slog.Error("Error connecting to port", "port", port, "error", err)
		os.Exit(1)
	}

}
