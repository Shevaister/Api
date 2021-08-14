package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

var (
	commentJSON = `{"ID":1,"PostID":1,"Name":"jinzu","Email":"jinzu@gmail.com","Body":"zzzzzz"}
`
	commentsJSON = `[{"ID":1,"PostID":1,"Name":"jinzu","Email":"jinzu@gmail.com","Body":"zzzzzz"},{"ID":2,"PostID":1,"Name":"jinzu","Email":"jinzu@gmail.com","Body":"zzzzzz"},{"ID":3,"PostID":1,"Name":"jinzu","Email":"jinzu@gmail.com","Body":"zzzzzz"}]
`
	createCommentJSON, _ = json.Marshal(map[string]interface{}{"id": 0, "PostID": 1, "Name": "jinzu", "email": "jinzu@gmail.com", "Body": "zzzzz"})
	updateCommentJSON, _ = json.Marshal(map[string]interface{}{"id": 0, "PostID": 1, "Name": "jinzu", "email": "jinzu@gmail.com", "Body": "vvvv"})
)

func TestReturnAllComments(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/comments")

	//Assertions
	if assert.NoError(t, ReturnAllComments(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, commentsJSON, rec.Body.String())
	}
}

func TestReturnComments(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/comments/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	//Assertions
	if assert.NoError(t, ReturnComment(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, commentJSON, rec.Body.String())
	}
}

func TestCreateComment(t *testing.T) {
	// Setup
	e := echo.New()

	handler := func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	}

	h := middleware.JWT([]byte("secret"))(handler)

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjI5MDEwNzgxLCJuYW1lIjoiZGVncmFlMDAwQGdtYWlsLmNvbSJ9.AKQMgE5pjeM6h4tKZoSU35YhthcHInoVlUucdIVfnG0"

	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(createCommentJSON))
	res := httptest.NewRecorder()
	req.Header.Set(echo.HeaderAuthorization, middleware.DefaultJWTConfig.AuthScheme+" "+token)
	c := e.NewContext(req, res)
	c.SetPath("/restricted/comments")
	assert.NoError(t, h(c))

	// Assertions
	if assert.NoError(t, CreateComment(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
	}
}

func TestUpdateComment(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(updateCommentJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/comments/:id")
	c.SetParamNames("id")
	c.SetParamValues("4")

	//Assertions
	if assert.NoError(t, UpdateComment(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeleteComment(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/comments/:id")
	c.SetParamNames("id")
	c.SetParamValues("4")

	//Assertions
	if assert.NoError(t, DeleteComment(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
