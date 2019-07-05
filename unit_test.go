package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rshindo/todo-go/common"
	"github.com/rshindo/todo-go/todo"
	"github.com/stretchr/testify/assert"
)

var app *gin.Engine

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestMain(m *testing.M) {
	db := common.Init()
	defer common.Close()
	app = NewApp(db)

	db.AutoMigrate(&todo.Todo{})

	code := m.Run()
	os.Exit(code)
}

func TestPing(t *testing.T) {

	w := performRequest(app, "GET", "/ping", nil)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"message":"pong"}`, w.Body.String())
}

func TestTodoCreate(t *testing.T) {

	w := performRequest(app, "POST", "/api/todo", strings.NewReader(`{"title":"Test","due_date":"2019-01-01"}`))

	assert.Equal(t, http.StatusCreated, w.Code)

	var response todo.TodoJSON
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	assert.Nil(t, err)
	assert.Equal(t, "Test", response.Title)
	assert.Equal(t, "2019-01-01", response.DueDate)

	id := response.ID
	idstr := fmt.Sprint(response.ID)
	w2 := performRequest(app, "GET", "/api/todo/"+idstr, nil)

	var response2 todo.TodoJSON
	err2 := json.Unmarshal([]byte(w2.Body.String()), &response2)
	assert.Nil(t, err2)
	assert.Equal(t, id, response2.ID)
	assert.Equal(t, "Testo", response2.Title)
	assert.Equal(t, "2019-01-01", response2.DueDate)
}
