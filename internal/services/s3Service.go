package services

import (
	"STDE_proj/internal/repositories"
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func PostFile(ctx context.Context, fileBytes []byte, objectKey string) (string, error) {
	return repositories.PostFile(ctx, fileBytes, objectKey)
}

func DeleteFile(ctx context.Context, objectKeys []string) error {
	return repositories.DeleteFile(ctx, objectKeys)
}

func GetFileURL(ctx context.Context, objectKeys string) string {
	return repositories.GetFileURL(objectKeys)
}

func DownloadFile(ctx context.Context, objectKey string) ([]byte, error) {
	return repositories.DownloadFile(ctx, objectKey)
}

func ListFiles(ctx context.Context, prefix string) (*s3.ListObjectsV2Output, error) {
	return repositories.ListFiles(prefix)
}

func ListBuckets(ctx context.Context) ([]types.Bucket, error) {
	return repositories.ListBuckets(ctx)
}
