package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io/ioutil"
	"net/http"
)

func main() {
	e := echo.New()
	//使用重定向中间件将http连接重定向到https
	e.Pre(middleware.HTTPSRedirect())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `
            <h1>Welcome to Echo!</h1>
            <h3>TLS certificates automatically installed from Let's Encrypt :)</h3>
        `)
	})

	e.POST("/", func(c echo.Context) error {
		r := c.Request() // c: echo.Context
		bodyReader := r.Body
		buf, err := ioutil.ReadAll(bodyReader)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(buf) // 二进制数据
		return c.HTML(http.StatusOK, `
            <h1>Welcome to Echo!</h1>
            <h3>TLS certificates automatically installed from Let's Encrypt :)</h3>
        `)
	})

	//go func() {
	//	e.Logger.Fatal(e.Start(":80"))
	//}()
	e.Logger.Fatal(e.StartTLS(":31140", "fullchain.pem", "privkey.pem"))
}