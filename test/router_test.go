package router_test

import (
	"github.com/gdsc-ys/fluentify-server/config"
	handler_test "github.com/gdsc-ys/fluentify-server/test/mocks/github.com/gdsc-ys/fluentify-server/src/handler"
	middleware_test "github.com/gdsc-ys/fluentify-server/test/mocks/github.com/gdsc-ys/fluentify-server/src/middleware"
	service_test "github.com/gdsc-ys/fluentify-server/test/mocks/github.com/gdsc-ys/fluentify-server/src/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gdsc-ys/fluentify-server/src/router"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	init := &config.Initialization{
		AuthMiddleware: middleware_test.NewMockAuthMiddleware(t),
		UserService:    service_test.NewMockUserService(t),
		StorageService: service_test.NewMockStorageService(t),
		TopicService:   service_test.NewMockTopicService(t),
		UserHandler:    handler_test.NewMockUserHandler(t),
		TopicHandler:   handler_test.NewMockTopicHandler(t),
	}
	e := router.Router(init)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello, World!", rec.Body.String())
}
