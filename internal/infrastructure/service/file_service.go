package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type FileService interface {
	UploadFile(file multipart.File, fileHeader *multipart.FileHeader, bucketName string) (string, error)
	DeleteFile(fileName, bucketName string) error
}

type fileService struct {
	minioClient *minio.Client
	endpoint    string
}

func NewFileService(endpoint, accessKey, secretKey string, useSSL bool) (FileService, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}

	return &fileService{
		minioClient: client,
		endpoint:    endpoint,
	}, nil
}

func (s *fileService) UploadFile(file multipart.File, fileHeader *multipart.FileHeader, bucketName string) (string, error) {
	ctx := context.Background()

	// Ensure bucket exists
	exists, err := s.minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		return "", err
	}
	if !exists {
		err = s.minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return "", err
		}
	}

	// Generate unique filename
	ext := filepath.Ext(fileHeader.Filename)
	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	// Upload file
	_, err = s.minioClient.PutObject(ctx, bucketName, fileName, file, fileHeader.Size, minio.PutObjectOptions{
		ContentType: fileHeader.Header.Get("Content-Type"),
	})
	if err != nil {
		return "", err
	}

	// Return file URL
	fileURL := fmt.Sprintf("http://%s/%s/%s", s.endpoint, bucketName, fileName)
	return fileURL, nil
}

func (s *fileService) DeleteFile(fileName, bucketName string) error {
	ctx := context.Background()
	return s.minioClient.RemoveObject(ctx, bucketName, fileName, minio.RemoveObjectOptions{})
}