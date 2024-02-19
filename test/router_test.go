package router_test

import (
	"github.com/gdsc-ys/fluentify-server/config"
	handler_test "github.com/gdsc-ys/fluentify-server/test/mocks/github.com/gdsc-ys/fluentify-server/src/handler"
	middleware_test "github.com/gdsc-ys/fluentify-server/test/mocks/github.com/gdsc-ys/fluentify-server/src/middleware"
	service_test "github.com/gdsc-ys/fluentify-server/test/mocks/github.com/gdsc-ys/fluentify-server/src/service"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gdsc-ys/fluentify-server/src/router"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	init := &config.Initialization{
		AuthMiddleware:  middleware_test.NewMockAuthMiddleware(t),
		UserService:     service_test.NewMockUserService(t),
		StorageService:  service_test.NewMockStorageService(t),
		TopicService:    service_test.NewMockTopicService(t),
		SentenceService: service_test.NewMockSentenceService(t),
		SceneService:    service_test.NewMockSceneService(t),
		UserHandler:     handler_test.NewMockUserHandler(t),
		TopicHandler:    handler_test.NewMockTopicHandler(t),
		SentenceHandler: handler_test.NewMockSentenceHandler(t),
		SceneHandler:    handler_test.NewMockSceneHandler(t),
		FeedbackHandler: handler_test.NewMockFeedbackHandler(t),
	}
	authMock := init.AuthMiddleware.(*middleware_test.MockAuthMiddleware)
	authMock.On("Verify").Return(echoMiddleware.Logger())
	e := router.Router(init)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello, World!", rec.Body.String())
}
