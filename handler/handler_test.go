package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"fmt"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestAddTask(t *testing.T) {
	// Setup
	e := echo.New()

	values := url.Values{}
	values.Set("title", "new-task")
	fmt.Println(values)

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
