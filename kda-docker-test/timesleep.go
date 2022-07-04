package main

import (
	"github.com/labstack/echo"
	"time"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		time.Sleep(time.Second * 5)
		return c.String(200, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1848"))
}
