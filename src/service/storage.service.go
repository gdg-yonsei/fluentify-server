package service

import (
	"bytes"
	"context"
	"firebase.google.com/go/v4/storage"
	"fmt"
	"github.com/gdsc-ys/fluentify-server/src/constant"
	"github.com/google/uuid"
	"io"
	"net/url"
)

type StorageService interface {
	UploadFile(file []byte, userId string) (string, error)
	GetFile(filePath string) ([]byte, error)
	GetFileUrl(filePath string) (string, error)
}

type StorageServiceImpl struct {
	storageClient *storage.Client
}

func (service *StorageServiceImpl) UploadFile(file []byte, userId string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), constant.StorageDefaultTimeout)
	defer cancel()

	bucket, err := service.storageClient.DefaultBucket()
	if err != nil {
		return "", err
	}

	fileName := uuid.New().String()
	path := fmt.Sprintf("%s%s/%s", constant.PublicBucketPath, userId, fileName)
	object := bucket.Object(path)
	writer := object.NewWriter(ctx)
	//Set the attribute
	//writer.ObjectAttrs.Metadata = map[string]string{"firebaseStorageDownloadTokens": userId}
	defer writer.Close()

	if _, err := io.Copy(writer, bytes.NewReader(file)); err != nil {
		return "", err
	}
	// Set access control if needed
	//if err := object.ACL().Set(context.Background(), storage.AllUsers, storage.RoleReader); err != nil {
	//	return "", err
	//}

	return path, nil
}

func (service *StorageServiceImpl) GetFile(filePath string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), constant.StorageDefaultTimeout)
	defer cancel()

	bucket, err := service.storageClient.DefaultBucket()
	if err != nil {
		return nil, err
	}

	object := bucket.Object(filePath)
	reader, err := object.NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(reader); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (service *StorageServiceImpl) GetFileUrl(filePath string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), constant.StorageDefaultTimeout)
	defer cancel()

	bucket, err := service.storageClient.DefaultBucket()
	if err != nil {
		return "", err
	}

	object := bucket.Object(filePath)
	attr, err := object.Attrs(ctx)
	if err != nil {
		return "", err
	}
	parsedName := url.QueryEscape(attr.Name)

	return fmt.Sprintf(
		"%s/b/%s/o/%s?alt=media",
		constant.StorageApiBaseUri,
		attr.Bucket,
		parsedName,
	), nil
}

func StorageServiceInit(storageClient *storage.Client) *StorageServiceImpl {
	return &StorageServiceImpl{
		storageClient: storageClient,
	}
}
