package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo"

	"github.com/subosito/gotenv"
)

type testTokenLoginResponseDTO struct {
	Token string `json:"token"`
}

type testLoginResponseDTO struct {
	Status  int                       `json:"status"`
	Success bool                      `json:"success"`
	Message string                    `json:"message"`
	Content testTokenLoginResponseDTO `json:"content"`
}

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

func TestLogin(t *testing.T) {

	gotenv.Load("../.env")

	e := echo.New()

	data := url.Values{}
	data.Set("username", "Marco")
	data.Set("password", "1234")

	req := httptest.NewRequest(http.MethodPost, "/user/login", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	if assert.NoError(t, Login(c)) {
		assert.Equal(t, http.StatusUnauthorized, res.Code)
	}

	data = url.Values{}
	data.Set("username", "Mario")
	data.Set("password", "123456")

	req = httptest.NewRequest(http.MethodPost, "/user/login", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	res = httptest.NewRecorder()

	c = e.NewContext(req, res)

	if assert.NoError(t, Login(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
	}

	var responseData testLoginResponseDTO
	json.NewDecoder(res.Result().Body).Decode(&responseData)

	bearer := "Bearer " + responseData.Content.Token

	req = httptest.NewRequest(http.MethodGet, "/restricted/user/duration", nil)
	req.Header.Set("Authorization", bearer)
	res = httptest.NewRecorder()

	c = e.NewContext(req, res)

	if assert.NoError(t, GetDurataionUsers(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
	}
}
