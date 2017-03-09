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
	e.GET("/todos/:id", getTask)
	e.POST("/todos/add", addTask)
	e.PUT("/todos/edit/:id", editTask)
	e.PUT("/todos/status/:id", changeTaskStatus)
	e.DELETE("/todos/:id", deleteTask)

	e.Logger.Fatal(e.Start(":9090"))
}

func getTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}
	return c.JSON(http.StatusOK, todoList[id])
}

func getList(c echo.Context) error {
	return c.JSON(http.StatusOK, todoList)
}

func addTask(c echo.Context) error {
	title := c.FormValue("title")
	if title == "" {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}
	todoList.Add(title)
	return c.JSON(http.StatusOK, "ok") // @todo まともなレスポンス返す
}

func changeTaskStatus(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}
	todoList.ChangeStatus(id)
	return c.JSON(http.StatusOK, "ok") // @todo まともなレスポンス返す
}

func editTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}
	title := c.FormValue("title")
	if title == "" {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	todoList.Edit(id, title)
	return c.JSON(http.StatusOK, "ok") // @todo まともなレスポンス返す
}

func deleteTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}
	todoList.Remove(id)
	return c.JSON(http.StatusOK, "ok") // @todo まともなレスポンス返す
}
