package storages

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/GroupProject3-Kelompok2/BE/app/config"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ClientUploader struct {
	cl         *storage.Client
	projectID  string
	bucketName string
	uploadPath string
}

func InitGCPClient() *storage.Client {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", config.GCP_CREDENTIAL)
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func UploadImage(c echo.Context, file *multipart.FileHeader) (string, error) {
	if file == nil {
		return "", nil
	}

	image, err := file.Open()
	if err != nil {
		return "", err
	}
	defer image.Close()

	sgcp := ClientUploader{
		cl:         InitGCPClient(),
		projectID:  config.GCP_PROJECTID,
		bucketName: config.GCP_BUCKETNAME,
		uploadPath: config.GCP_PATH,
	}

	imageURL, err := sgcp.UploadFile(image, file.Filename)
	if err != nil {
		return "", err
	}

	return imageURL, nil
}

// UploadFile uploads an object
func (c *ClientUploader) UploadFile(file multipart.File, object string) (string, error) {
	rand := uuid.New().String()
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := c.cl.Bucket(c.bucketName).Object(c.uploadPath + object + rand).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return "", fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("Writer.Close: %v", err)
	}

	escapedObject := strings.ReplaceAll(object, " ", "%20")
	fileURL := "https://storage.googleapis.com/" + c.bucketName + "/" + c.uploadPath + escapedObject + rand
	return fileURL, nil
}
