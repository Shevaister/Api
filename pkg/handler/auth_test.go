package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	testUser, _ = json.Marshal(map[string]interface{}{"email": "ddg@gmail.com", "password": "ghgwg"})
)

func TestGoogleSignIn(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/signin/google")

	//Assertions
	if assert.NoError(t, GoogleSignIn(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestFacebookSignIn(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/signin/facebook")

	//Assertions
	if assert.NoError(t, GoogleSignIn(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestSignUp(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(testUser))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/signup")

	//Assertions
	if assert.NoError(t, SignUp(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestSignIn(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(testUser))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/signin")

	//Assertions
	if assert.NoError(t, SignIn(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
