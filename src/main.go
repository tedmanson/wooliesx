package main

import (
	"net/http"
	"sort"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tedmanson/wool/src/wooliesx"
)

func main() {

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.GET("/user", getUser)

	e.Use(userTokenAuthMiddleware)
	e.Use(wooliesSDKMiddleware(wooliesx.New("http://dev-wooliesx-recruitment.azurewebsites.net/api/")))

	e.GET("/products", getProducts)

	e.Logger.Fatal(e.Start(":1323"))

	// s := &http.Server{
	// 	Addr:           ":1323",
	// 	Handler:        e,
	// 	ReadTimeout:    5,
	// 	WriteTimeout:   5,
	// 	MaxHeaderBytes: 1 << 20,
	// }

	// e.Logger.Fatal(e.StartServer(s))
}

// e.GET("/users/:id", getUser)
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
	var token = c.QueryParam("token")
	var sortOption = c.QueryParam("sortOption")

	wx := c.Get("wooliesx").(*wooliesx.SDK)
	if wx == nil {
		echo.NewHTTPError(http.StatusFailedDependency, "Product listing currently unavailable")
	}

	products, err := wx.GetProducts(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Product listing currently unavailable")
	}

	switch strings.ToLower(sortOption) {
	case "ascending":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Name < products[j].Name
		})
		break

	case "descending":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Name > products[j].Name
		})
		break

	case "low":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Price < products[j].Price
		})
		break

	case "high":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Price > products[j].Price
		})
		break

	case "recommended":
		//TODO "Recommended" - this will call the "shopperHistory" resource to get a list of customers orders and needs to return based on popularity,
		break
	}

	return c.JSON(http.StatusOK, products)
}

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
