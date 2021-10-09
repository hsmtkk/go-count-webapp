package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	cnt := &counter{count: 0}
	// Routes
	e.GET("/", hello)
	e.GET("/count", cnt.countup)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

type counter struct {
	count int
}

// /count にアクセスすると、これまでのアクセス数を表示する
func (c *counter) countup(ctx echo.Context) error {
	c.count += 1
	return ctx.String(http.StatusOK, strconv.Itoa(c.count))
}
