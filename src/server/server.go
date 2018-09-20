package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Run loads and starts the echo http server
func Run() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Pre(middleware.RemoveTrailingSlash())

	e.GET("/user", getUser)

	e.Use(userTokenAuthMiddleware)
	e.Use(wooliesSDKMiddleware)

	e.GET("/sort", getProducts)
	e.GET("/trolleyCalculator", getTrolley)

	e.Logger.Fatal(e.Start(":8080"))
}
