package handler

import (
	"context"
	"errors"
	pb "github.com/gdsc-ys/fluentify-server/gen/proto"
	"github.com/gdsc-ys/fluentify-server/src/constant"
	"github.com/gdsc-ys/fluentify-server/src/model"
	"github.com/gdsc-ys/fluentify-server/src/service"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io/fs"
	"log"
	"math"
	"net/http"
	"os"
	"time"
)

type FeedbackHandler interface {
	GetPronunciationFeedback(c echo.Context) error
	GetCommunicationFeedback(c echo.Context) error
}

type FeedbackHandlerImpl struct {
	storageService service.StorageService
}

func (handler *FeedbackHandlerImpl) GetPronunciationFeedback(c echo.Context) error {
	var request = pb.GetPronunciationFeedbackRequest{}
	if err := c.Bind(&request); err != nil {
		return model.NewCustomHTTPError(http.StatusBadRequest, err)
	}

	hardCodedAudioFile := "example1.m4a"
	sharedFilePath := constant.SharedAudioPath + "/" + hardCodedAudioFile
	if !existsFile(sharedFilePath) {
		bucketPath := "audio/" + hardCodedAudioFile
		bytes, err := handler.storageService.GetFile(bucketPath)
		if err != nil {
			return model.NewCustomHTTPError(http.StatusInternalServerError, err)
		}

		err = writeFile(sharedFilePath, bytes)
		if err != nil {
			return model.NewCustomHTTPError(http.StatusInternalServerError, err)
		}
	}

	conn, err := grpc.Dial(os.Getenv("AI_SERVER_HOST"), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Printf("error connecting to grpc server: %v", err)
		return model.NewCustomHTTPError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
	defer conn.Close()
	//client 생성
	client := pb.NewPronunciationFeedbackServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	grpcRequest := pb.PronunciationFeedbackRequest{
		Sentence:  "It's autumn now, and the leaves are turning beautiful colors.",
		AudioPath: hardCodedAudioFile,
		Tip:       "Say 'aw-tum,' not 'ay-tum.'",
	}

	response, err := client.PronunciationFeedback(ctx, &grpcRequest)
	if err != nil {
		log.Printf("grpc request failed: %v", err)
		return model.NewCustomHTTPError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	pronunciationScore := int32(math.Round(response.GetPronunciationScore() * 100))
	decibel := int32(math.Round(response.GetDecibel()))
	speechRate := int32(math.Round(response.GetSpeechRate()))

	result := pb.GetPronunciationFeedbackResponse{
		PronunciationFeedback: &pb.PronunciationFeedbackDTO{
			SentenceId:         request.GetSentenceId(),
			IncorrectIndexes:   response.WrongIdxMajor,
			PronunciationScore: pronunciationScore,
			VolumeScore:        decibel,
			SpeedScore:         speechRate,
			OverallFeedback:    response.GetPositiveFeedback(),
		},
	}

	return c.JSON(http.StatusOK, &result)
}

func (handler *FeedbackHandlerImpl) GetCommunicationFeedback(c echo.Context) error {
	var request = pb.GetCommunicationFeedbackRequest{}
	if err := c.Bind(&request); err != nil {
		return model.NewCustomHTTPError(http.StatusBadRequest, err)
	}

	hardCodedAudioFile := "example1.m4a"
	sharedFilePath := constant.SharedAudioPath + "/" + hardCodedAudioFile
	if !existsFile(sharedFilePath) {
		bucketPath := "audio/" + hardCodedAudioFile
		bytes, err := handler.storageService.GetFile(bucketPath)
		if err != nil {
			return model.NewCustomHTTPError(http.StatusInternalServerError, err)
		}

		err = writeFile(sharedFilePath, bytes)
		if err != nil {
			return model.NewCustomHTTPError(http.StatusInternalServerError, err)
		}
	}

	conn, err := grpc.Dial(os.Getenv("AI_SERVER_HOST"), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Printf("error connecting to grpc server: %v", err)
		return model.NewCustomHTTPError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))

	}
	defer conn.Close()

	//client 생성
	client := pb.NewCommunicationFeedbackServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	grpcRequest := pb.CommunicationFeedbackRequest{
		Context:        "Let's imagine that you are a brave captain of a big ship. You are sailing on the high seas. Suddenly, you see a beautiful sunset. Look at this picture and tell me...",
		Question:       "What colors can you see in the sky?",
		ExpectedAnswer: "I can see red, orange, and yellow.",
		AudioPath:      "example1.m4a",
		ImgPath:        "img/1070.jpg",
	}

	response, err := client.CommunicationFeedback(ctx, &grpcRequest)
	if err != nil {
		log.Printf("grpc request failed: %v", err)
		return model.NewCustomHTTPError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	result := pb.GetCommunicationFeedbackResponse{
		CommunicationFeedback: &pb.CommunicationFeedbackDTO{
			SceneId:         request.GetSceneId(),
			OverallFeedback: response.GetPositiveFeedback(),
		},
	}

	return c.JSON(http.StatusOK, &result)
}

func existsFile(fileName string) bool {
	_, err := os.Stat(fileName)

	if errors.Is(err, fs.ErrNotExist) {
		return false
	} else {
		return true
	}
}

func writeFile(fileName string, bytes []byte) error {
	fileToWrite, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer fileToWrite.Close()

	_, err = fileToWrite.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

func FeedbackHandlerInit(storageService service.StorageService) *FeedbackHandlerImpl {
	return &FeedbackHandlerImpl{
		storageService: storageService,
	}
}
