package rest

import (
	"github.com/kataras/iris/v12/httptest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_invalid_username(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.SetBasicAuth("dummyAdmin", "D3v1ceMgr")
	basicAuthHandler := newBasicAuthHandler(configForTesting.UserName, configForTesting.Password)

	httptest.Do(rec, req, basicAuthHandler)
	assert.Equal(t, http.StatusUnauthorized, rec.Result().StatusCode)
}

func Test_invalid_password(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.SetBasicAuth("admin", "dummyPassword")
	basicAuthHandler := newBasicAuthHandler(configForTesting.UserName, configForTesting.Password)

	httptest.Do(rec, req, basicAuthHandler)
	assert.Equal(t, http.StatusUnauthorized, rec.Result().StatusCode)
}

func Test_valid_credentials(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.SetBasicAuth("admin", "D3v1ceMgr")
	basicAuthHandler := newBasicAuthHandler(configForTesting.UserName, configForTesting.Password)

	httptest.Do(rec, req, basicAuthHandler)
	assert.Equal(t, http.StatusOK, rec.Result().StatusCode)
}
