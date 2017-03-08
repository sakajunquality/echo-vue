package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sakajunquality/echo-vue/todo"
)

var todoList todo.List

func main() {
	todoList = todo.List{}

	// Add some dummy daya
	todoList.Add("test1").Add("test2").Add("test3").Remove(2).Add("test4")

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.GET("/todos", getList)
	e.POST("/todos/add", addTask)
	e.DELETE("/todos/:id", deleteTask)

	e.Logger.Fatal(e.Start(":9090"))
}

func getList(c echo.Context) error {
	return c.JSON(http.StatusOK, todoList)
}

func addTask(c echo.Context) error {
	title := c.FormValue("title")
	// @todo titleのチェック
	todoList.Add(title)
	return c.JSON(http.StatusOK, "ok") // @todo まともなレスポンス返す
}

func deleteTask(c echo.Context) error {
	// @todo intで取れてるかのチェック
	id, _ := strconv.Atoi(c.Param("id"))
	todoList.Remove(id)
	return c.JSON(http.StatusOK, "ok") // @todo まともなレスポンス返す
}
