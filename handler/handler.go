package handler

import (
	"net/http"
	"strconv"

	"fmt"

	"github.com/labstack/echo"
	"github.com/sakajunquality/echo-vue/todo"
)

var todoList todo.List

type resp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Handler struct {
	List todo.List
}

func (h *Handler) Init() {
	todoList = todo.List{}
	todoList.Add("test1").Add("test2").Add("test3").Remove(2).Add("test4")
}

func (h *Handler) GetTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			resp{Code: http.StatusBadRequest, Message: "Invalid ID specification"})
	}
	td, ok := todoList[id]
	if !ok {
		return c.JSON(http.StatusNotFound,
			resp{
				Code:    http.StatusBadRequest,
				Message: fmt.Sprintf("ID: %s was not found", strconv.Itoa(id)),
			})
	}

	return c.JSON(http.StatusOK, td)
}

func (h *Handler) GetList(c echo.Context) error {
	return c.JSON(http.StatusOK, todoList)
}

func (h *Handler) AddTask(c echo.Context) error {
	title := c.FormValue("title")
	if title == "" {
		return c.JSON(http.StatusBadRequest,
			resp{Code: http.StatusBadRequest, Message: "Title is not valid"})
	}
	todoList.Add(title)

	return c.JSON(http.StatusOK,
		resp{
			Code:    http.StatusOK,
			Message: fmt.Sprintf("New Task %s was added", title),
		})
}

func (h *Handler) ChangeTaskStatus(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			resp{Code: http.StatusBadRequest, Message: "Invalid ID specification"})
	}
	todoList.ChangeStatus(id)
	return c.JSON(http.StatusOK,
		resp{
			Code:    http.StatusOK,
			Message: fmt.Sprintf("Status of task #%s was changed", strconv.Itoa(id)),
		})
}

func (h *Handler) EditTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			resp{Code: http.StatusBadRequest, Message: "Invalid ID specification"})
	}
	title := c.FormValue("title")
	if title == "" {
		return c.JSON(http.StatusBadRequest,
			resp{Code: http.StatusBadRequest, Message: "Title is not valid"})
	}

	todoList.Edit(id, title)
	return c.JSON(http.StatusOK,
		resp{
			Code:    http.StatusOK,
			Message: fmt.Sprintf("Title of task #%s was changed", strconv.Itoa(id)),
		})
}

func (h *Handler) DeleteTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			resp{Code: http.StatusBadRequest, Message: "Invalid ID specification"})
	}
	todoList.Remove(id)
	return c.JSON(http.StatusOK,
		resp{
			Code:    http.StatusOK,
			Message: fmt.Sprintf("Task #%s was deleted", strconv.Itoa(id)),
		})
}
