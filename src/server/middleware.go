package server

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/tedmanson/wooliesx/src/wooliesx"
)

func userTokenAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if strings.HasSuffix(c.Request().URL.Path, "user") {
			var user = struct {
				User  string `json:"user"`
				Token string `json:"token"`
			}{
				User:  "Scott Tedmanson",
				Token: "c722c138-b07f-4ee5-8c6f-7d603d552479",
			}

			return c.JSON(http.StatusOK, user)
		}

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
