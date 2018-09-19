package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/tedmanson/wool/src/wooliesx"
)

func userTokenAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.QueryParam("token") == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
		}
		return next(c)
	}
}

func wooliesSDKMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var wx = wooliesx.New("http://dev-wooliesx-recruitment.azurewebsites.net/api/", c.QueryParam("token"))
		c.Set("wooliesx", wx)
		return next(c)
	}
}
