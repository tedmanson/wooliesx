package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tedmanson/wool/src/wooliesx"
)

// Run loads and starts the echo http server
func Run() {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.GET("/user", getUser)

	e.Use(userTokenAuthMiddleware)
	e.Use(wooliesSDKMiddleware(wooliesx.New("http://dev-wooliesx-recruitment.azurewebsites.net/api/")))

	e.GET("/sort", getProducts)

	e.Logger.Fatal(e.Start(":8080"))
}
