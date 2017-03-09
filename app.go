package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sakajunquality/echo-vue/handler"
)

func main() {
	e := echo.New()
	h := handler.Handler{}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// テストデータとりあえず入れる
	h.Init()

	e.GET("/todos", h.GetList)
	e.GET("/todos/:id", h.GetTask)
	e.POST("/todos/add", h.AddTask)
	e.PUT("/todos/edit/:id", h.EditTask)
	e.PUT("/todos/status/:id", h.ChangeTaskStatus)
	e.DELETE("/todos/:id", h.DeleteTask)

	e.Logger.Fatal(e.Start(":9090"))
}
