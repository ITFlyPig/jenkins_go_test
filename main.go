package main

import "github.com/labstack/echo/v4"

func main() {
	e := echo.New()
	// 配置路由
	e.GET("/hellow", func(c echo.Context) error {
		//return c.JSON(100, "hello")
		return c.String(200, "hello")
	})
	// 启动Web服务
	e.Start(":8090")
}
