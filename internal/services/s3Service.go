package services

import (
	"STDE_proj/internal/repositories"
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"mime/multipart"
)

func PostFile(ctx context.Context, file multipart.File, fileName string) (string, error) {
	return repositories.PostFile(file, fileName)
}

func DeleteFile(ctx context.Context, fileName string) error {
	return repositories.DeleteFile(fileName)
}

func GetFileURL(ctx context.Context, fileName string) string {
	return repositories.GetFileURL(fileName)
}

func DownloadFile(ctx context.Context, fileName string) ([]byte, error) {
	return repositories.DownloadFile(fileName)
}

func ListFiles(ctx context.Context, prefix string) (*s3.ListObjectsV2Output, error) {
	return repositories.ListFiles(prefix)
}

func ListBuckets(ctx context.Context) (*s3.ListBucketsOutput, error) {
	return repositories.ListBuckets()
}
