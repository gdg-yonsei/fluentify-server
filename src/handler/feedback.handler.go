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
	"strings"
	"unicode"
	"unicode/utf8"
)

type FeedbackHandler interface {
	GetPronunciationFeedback(c echo.Context) error
	GetCommunicationFeedback(c echo.Context) error
}

type FeedbackHandlerImpl struct {
	sentenceService service.SentenceService
	sceneService    service.SceneService
	storageService  service.StorageService
}

func (handler *FeedbackHandlerImpl) GetPronunciationFeedback(c echo.Context) error {
	var request = pb.GetPronunciationFeedbackRequest{}
	if err := c.Bind(&request); err != nil {
		return model.NewCustomHTTPError(http.StatusBadRequest, err)
	}
	if request.GetSentenceId() == "" {
		return model.NewCustomHTTPError(http.StatusBadRequest, "sentence_id is required")
	}
	sentence, err := handler.sentenceService.GetSentence(request.GetSentenceId())
	if err != nil {
		return model.NewCustomHTTPError(http.StatusInternalServerError, err)
	}

	bucketPath := request.GetAudioFileUrl()
	splitPath := strings.Split(bucketPath, "/")
	fileName := splitPath[len(splitPath)-1]
	sharedFilePath := constant.SharedAudioPath + "/" + fileName
	if !existsFile(sharedFilePath) {
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

	ctx, cancel := context.WithTimeout(context.Background(), constant.GrpcDefaultTimeout)
	defer cancel()

	grpcRequest := pb.PronunciationFeedbackRequest{
		Sentence:  sentence.Text,
		AudioPath: fileName,
		Tip:       sentence.Tip,
	}

	feedbackResponse, err := client.PronunciationFeedback(ctx, &grpcRequest)
	if err != nil {
		log.Printf("grpc request failed: %v", err)
		return model.NewCustomHTTPError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	pronunciationScore := int32(math.Round(feedbackResponse.GetPronunciationScore()))
	decibel := int32(math.Round(feedbackResponse.GetDecibel()))
	speechRate := int32(math.Round(feedbackResponse.GetSpeechRate()))
	transcript := firstToUpper(feedbackResponse.GetTranscript())

	result := pb.GetPronunciationFeedbackResponse{
		PronunciationFeedback: &pb.PronunciationFeedbackDTO{
			SentenceId:         request.GetSentenceId(),
			Transcript:         transcript,
			IncorrectIndexes:   feedbackResponse.GetIncorrectIndexes(),
			PronunciationScore: pronunciationScore,
			VolumeScore:        decibel,
			SpeedScore:         speechRate,
			OverallFeedback:    feedbackResponse.GetPositiveFeedback(),
		},
	}

	return c.JSON(http.StatusOK, &result)
}

func (handler *FeedbackHandlerImpl) GetCommunicationFeedback(c echo.Context) error {
	var request = pb.GetCommunicationFeedbackRequest{}
	if err := c.Bind(&request); err != nil {
		return model.NewCustomHTTPError(http.StatusBadRequest, err)
	}
	if request.GetSceneId() == "" {
		return model.NewCustomHTTPError(http.StatusBadRequest, "scene_id is required")
	}
	scene, err := handler.sceneService.GetScene(request.GetSceneId())
	if err != nil {
		return model.NewCustomHTTPError(http.StatusInternalServerError, err)
	}

	bucketPath := request.GetAudioFileUrl()
	splitPath := strings.Split(bucketPath, "/")
	fileName := splitPath[len(splitPath)-1]
	sharedFilePath := constant.SharedAudioPath + "/" + fileName
	if !existsFile(sharedFilePath) {
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

	ctx, cancel := context.WithTimeout(context.Background(), constant.GrpcDefaultTimeout)
	defer cancel()

	grpcRequest := pb.CommunicationFeedbackRequest{
		Context:        scene.Context,
		Question:       scene.Question,
		ExpectedAnswer: scene.ExpectedAnswer,
		AudioPath:      fileName,
		ImgPath:        scene.ImageUrl,
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
	err := os.MkdirAll(constant.SharedAudioPath, os.ModePerm)
	if err != nil {
		return err
	}

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

func firstToUpper(s string) string {
	r, size := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError && size <= 1 {
		return s
	}
	lowerCase := strings.ToLower(s)
	return string(unicode.ToUpper(r)) + lowerCase[size:]
}

func FeedbackHandlerInit(sentenceService service.SentenceService, sceneService service.SceneService, storageService service.StorageService) *FeedbackHandlerImpl {
	return &FeedbackHandlerImpl{
		sentenceService: sentenceService,
		sceneService:    sceneService,
		storageService:  storageService,
	}
}
