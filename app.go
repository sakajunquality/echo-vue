package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sakajunquality/echo-vue/todo"
)

var todoList todo.List

func main() {
	todoList = todo.List{}

	// Add some dummy daya
	todoList.Add("test1").Add("test2").Add("test3").Remove(2).Add("test4")

	e := echo.New()
	e.GET("/", getList)
	e.POST("/add", addTask)

	e.Logger.Fatal(e.Start(":9090"))
}

func getList(c echo.Context) error {
	return c.JSON(http.StatusOK, todoList)
}

func addTask(c echo.Context) error {
	title := c.FormValue("title")
	todoList.Add(title)
	return c.JSON(http.StatusOK, "ok")
}
