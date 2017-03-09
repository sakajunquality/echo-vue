package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var initialListJSON = `{"1":{"title":"test1","completed":false},"3":{"title":"test3","completed":false},"4":{"title":"test4","completed":false}}`

func TestAddTask(t *testing.T) {
	e := echo.New()

	values := url.Values{}
	values.Set("title", "new-task")

	req, err := http.NewRequest(echo.POST, "/todos/add", strings.NewReader(values.Encode()))
	if assert.NoError(t, err) {
		// ホントは、echo.MIMEApplicationJSONにする
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &Handler{}
		h.Init()

		// Assertions
		if assert.NoError(t, h.AddTask(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	}
}

func TestAddInvalidTask(t *testing.T) {
	e := echo.New()

	values := url.Values{}
	values.Set("title", "") //empty

	req, err := http.NewRequest(echo.POST, "/todos/add", strings.NewReader(values.Encode()))
	if assert.NoError(t, err) {
		// ホントは、echo.MIMEApplicationJSONにする
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := &Handler{}
		h.Init()

		// Assertions
		if assert.NoError(t, h.AddTask(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	}
}

func TestGetTaskList(t *testing.T) {
	e := echo.New()
	req := new(http.Request)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/todos")
	h := &Handler{}
	h.Init()

	if assert.NoError(t, h.GetList(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, initialListJSON, rec.Body.String())
	}
}

func TestGetTask(t *testing.T) {
	e := echo.New()
	req := new(http.Request)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/todos/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	h := &Handler{}
	h.Init()

	if assert.NoError(t, h.GetTask(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetInvalidTask(t *testing.T) {
	e := echo.New()
	req := new(http.Request)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/todos/:id")
	c.SetParamNames("id")
	c.SetParamValues("777")

	h := &Handler{}
	h.Init()

	if assert.NoError(t, h.GetTask(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}

func TestEditTask(t *testing.T) {
	e := echo.New()

	values := url.Values{}
	values.Set("title", "yet another task name")

	req, err := http.NewRequest(echo.PUT, "", strings.NewReader(values.Encode()))
	if assert.NoError(t, err) {
		// ホントは、echo.MIMEApplicationJSONにする
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/todos/edit/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		h := &Handler{}
		h.Init()

		// Assertions
		if assert.NoError(t, h.EditTask(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	}
}

func TestChangeTaskStatus(t *testing.T) {
	e := echo.New()

	req, err := http.NewRequest(echo.PUT, "", nil)
	if assert.NoError(t, err) {
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/todos/status/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		h := &Handler{}
		h.Init()

		// Assertions
		if assert.NoError(t, h.ChangeTaskStatus(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	}
}

func TestDeleteTask(t *testing.T) {
	e := echo.New()

	req, err := http.NewRequest(echo.DELETE, "", nil)
	if assert.NoError(t, err) {
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/todos/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		h := &Handler{}
		h.Init()

		// Assertions
		if assert.NoError(t, h.DeleteTask(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	}
}
