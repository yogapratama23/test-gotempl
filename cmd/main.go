package main

import (
	"github.com/labstack/echo/v4"
	"github.com/yogapratama23/test-gotempl/handlers"
)

func main() {
	app := echo.New()

	userHandler := handlers.UserHandler{}
	app.GET("/user", userHandler.HandleUserShow)

	app.Start(":3030")
}
