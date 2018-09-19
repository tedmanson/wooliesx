package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/tedmanson/wool/src/wooliesx"
)

func getUser(c echo.Context) error {
	var user = struct {
		User  string `json:"user"`
		Token string `json:"token"`
	}{
		User:  "Scott Tedmanson",
		Token: "c722c138-b07f-4ee5-8c6f-7d603d552479",
	}

	return c.JSON(http.StatusOK, user)
}

func getProducts(c echo.Context) error {
	var sortOption = c.QueryParam("sortOption")

	wx := c.Get("wooliesx").(*wooliesx.SDK)
	if wx == nil {
		echo.NewHTTPError(http.StatusFailedDependency, "Product listing currently unavailable")
	}

	products, err := wx.GetProducts()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Product listing currently unavailable")
	}

	if sortOption != "" {
		products = wx.SortProducts(products, sortOption)
	}

	return c.JSON(http.StatusOK, products)
}
