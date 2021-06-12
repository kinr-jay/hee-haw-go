package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"handlers"
)

func main() {
	e := echo.New()
	
	e.GET("/", func(c echo.Context) error {
		handlers.Hello()
		return c.String(http.StatusOK, "Yeeeee-haw!")
	})

	e.Logger.Fatal(e.Start(":8000"))
}
