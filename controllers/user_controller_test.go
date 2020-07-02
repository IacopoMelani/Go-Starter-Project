package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/jmoiron/sqlx"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"

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

	if err := gotenv.Load("./../.env"); err != nil {
		t.Fatal("Errore caricamento configurazione")
	}
	db.InitConnection(os.Getenv("SQL_DRIVER"), os.Getenv("STRING_CONNECTION"))

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/user/all", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, GetAllUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

	db := db.GetConnection().(*sqlx.DB)

	db.Close()

	req = httptest.NewRequest(http.MethodGet, "/user/all", nil)
	rec = httptest.NewRecorder()

	c = e.NewContext(req, rec)

	if assert.NoError(t, GetAllUser(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	}
}

func TestGetDurataionUsers(t *testing.T) {

	if err := gotenv.Load("../.env"); err != nil {
		t.Fatal("Errore caricamento configurazione")
	}

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/user/duration", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, GetDurataionUsers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestLogin(t *testing.T) {

	if err := gotenv.Load("../.env"); err != nil {
		t.Fatal("Errore caricamento configurazione")
	}

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
	if err := json.NewDecoder(res.Result().Body).Decode(&responseData); err != nil {
		t.Error(err.Error())
	}

	bearer := "Bearer " + responseData.Content.Token

	req = httptest.NewRequest(http.MethodGet, "/restricted/user/duration", nil)
	req.Header.Set("Authorization", bearer)
	res = httptest.NewRecorder()

	c = e.NewContext(req, res)

	if assert.NoError(t, GetDurataionUsers(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
	}
}
