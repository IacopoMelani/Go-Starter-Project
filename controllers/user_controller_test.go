package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo"

	"github.com/subosito/gotenv"
)

func TestGetAllUser(t *testing.T) {

	gotenv.Load("../.env")

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/user/all", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, GetAllUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetDurataionUsers(t *testing.T) {

	gotenv.Load("../.env")

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/user/duration", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, GetDurataionUsers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
