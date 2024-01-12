package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gdsc-ys/fluentify-server/src/router"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	e := router.Router()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello, World!", rec.Body.String())
}
