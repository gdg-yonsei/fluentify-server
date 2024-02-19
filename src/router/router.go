package router

import (
	"context"
	"github.com/gdsc-ys/fluentify-server/config"
	pb "github.com/gdsc-ys/fluentify-server/gen/proto"
	"github.com/gdsc-ys/fluentify-server/src/middleware"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func Router(init *config.Initialization) *echo.Echo {
	e := echo.New()

	e.Debug = true

	// Middleware
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())
	e.HTTPErrorHandler = middleware.CustomHTTPErrorHandler

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/GetUser", init.UserHandler.GetUser)
	e.POST("/UpdateUser", init.UserHandler.UpdateUser)
	e.POST("/DeleteUser", init.UserHandler.DeleteUser)

	e.POST("/ListTopics", init.TopicHandler.ListTopics)
	e.POST("/GetTopic", init.TopicHandler.GetTopic)

	e.POST("/GetSentence", init.SentenceHandler.GetSentence)
	e.POST("/GetScene", init.SceneHandler.GetScene)

	e.GET("/PingHello", func(c echo.Context) error {
		name := c.Param("name")
		conn, err := grpc.Dial(os.Getenv("AI_SERVER_HOST"), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect grpc: %v", err)
		}
		defer conn.Close()
		//client 생성
		client := pb.NewHelloServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		//Ping 전송
		response, err := client.Hello(ctx, &pb.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("could not request grpc: %v", err)
		}

		return c.String(http.StatusOK, response.GetMessage())
	})

	return e
}
