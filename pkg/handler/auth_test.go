package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
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
