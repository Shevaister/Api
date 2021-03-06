package handler

import (
	_ "Api/config"
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
	stringPostsJSON = `[{"ID":1,"UserID":1,"Title":"not","Body":"dsadasdadwaa"},{"ID":2,"UserID":1,"Title":"not","Body":"dsadasdadwaa"}]
`
	stringPostJSON = `{"ID":1,"UserID":1,"Title":"not","Body":"dsadasdadwaa"}
`
	createPostJSON, _ = json.Marshal(map[string]interface{}{"id": 0, "UserID": 1, "Title": "not", "Body": "x"})
	updatePostJSON, _ = json.Marshal(map[string]interface{}{"id": 0, "UserID": 5, "Title": "not", "Body": "y"})
)

func TestReturnAllPosts(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/posts")

	//Assertions
	if assert.NoError(t, ReturnAllPosts(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, stringPostsJSON, rec.Body.String())
	}
}

func TestReturnPost(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/posts/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	//Assertions
	if assert.NoError(t, ReturnPost(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, stringPostJSON, rec.Body.String())
	}
}

func TestCreatePost(t *testing.T) {
	// Setup
	e := echo.New()

	handler := func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	}

	h := middleware.JWT([]byte("secret"))(handler)

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjI5NDg1NTI3LCJuYW1lIjoieWV2aGVuLnNoZXZjaHVrQG51cmUudWEifQ.elg5kf-2Pt2mYl9e5V75NF1IAWh2ZCKcYS134Sj20Ro"

	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(createPostJSON))
	res := httptest.NewRecorder()
	req.Header.Set(echo.HeaderAuthorization, middleware.DefaultJWTConfig.AuthScheme+" "+token)
	c := e.NewContext(req, res)
	c.SetPath("/restricted/posts")
	assert.NoError(t, h(c))

	// Assertions
	if assert.NoError(t, CreatePost(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
	}
}

func TestUpdatePost(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(updatePostJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/posts/:id")
	c.SetParamNames("id")
	c.SetParamValues("3")

	//Assertions
	if assert.NoError(t, UpdatePost(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeletePost(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/posts/:id")
	c.SetParamNames("id")
	c.SetParamValues("3")

	//Assertions
	if assert.NoError(t, DeletePost(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
