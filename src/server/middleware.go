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

func wooliesSDKMiddleware(wooliesx *wooliesx.SDK) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("wooliesx", wooliesx)
			return next(c)
		}
	}
}
