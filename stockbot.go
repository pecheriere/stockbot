package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/pecheriere/stockbot/apiai"
	"github.com/labstack/echo/middleware"
)

func main() {
	config := Configuration()

	e := echo.New()

	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (error, bool) {
		if username == config.Auth.Username && password == config.Auth.Password {
			fmt.Println("Is Authorized")
			return nil, true
		}
		return nil, false
	}))

	e.POST("/fulfill", func(c echo.Context) error {

		fmt.Println("Received a fulfillment request")

		fr := new(apiai.FulfillmentRequest)
		if err := c.Bind(fr); err != nil {
			return err
		}

		err, response := apiai.Fulfill(fr)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, response)

	})

	e.GET("/", func(c echo.Context) error {
		fmt.Println("Received a GET request")
		return c.JSON(http.StatusOK, "hello")
	})

	e.Logger.Fatal(e.Start(":3253"))
}