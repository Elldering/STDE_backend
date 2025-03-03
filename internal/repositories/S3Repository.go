package repositories

import (
	"STDE_proj/configs"
	"STDE_proj/utils/hash"
	"STDE_proj/utils/time_web_s3"
	"bytes"
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
	"log"
	"strings"
	"time"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
	"net/url"
)

func PostFile(ctx context.Context, fileBytes []byte, objectKey string) (string, error) {
	reader := bytes.NewReader(fileBytes)
	objectKey = strings.TrimLeft(objectKey, "/")

	_, err := time_web_s3.S3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:         aws.String(configs.AppConfig.S3.Bucket),
		Key:            aws.String(objectKey),
		Body:           reader,
		ChecksumSHA256: aws.String(hash.ComputeSHA256Hex(fileBytes)),
		ContentType:    aws.String("image/png"),
	})
	if err != nil {
		return "", fmt.Errorf("ошибка загрузки файла: %w", err)
	}

	encodedKey := url.PathEscape(objectKey)
	fileURL := fmt.Sprintf("%s/%s/%s", configs.AppConfig.S3.Endpoint, configs.AppConfig.S3.Bucket, encodedKey)
	return fileURL, nil
}

func DeleteFile(ctx context.Context, objectKeys []string) error {
	var objectIds []types.ObjectIdentifier
	for _, key := range objectKeys {
		key = strings.TrimLeft(key, "/") // Убираем ведущие слеши
		objectIds = append(objectIds, types.ObjectIdentifier{Key: aws.String(key)})
	}

	output, err := time_web_s3.S3Client.DeleteObjects(ctx, &s3.DeleteObjectsInput{
		Bucket: aws.String(configs.AppConfig.S3.Bucket),
		Delete: &types.Delete{Objects: objectIds, Quiet: aws.Bool(true)},
	})
	if err != nil || len(output.Errors) > 0 {
		log.Printf("Ошибка удаления объектов из бакета %s.\n", configs.AppConfig.S3.Bucket)
		if err != nil {
			var noBucket *types.NoSuchBucket
			if errors.As(err, &noBucket) {
				log.Printf("Бакет %s не существует.\n", configs.AppConfig.S3.Bucket)
				return noBucket
			}
			return err
		}
		if len(output.Errors) > 0 {
			for _, outErr := range output.Errors {
				log.Printf("%s: %s\n", *outErr.Key, *outErr.Message)
			}
			return fmt.Errorf("%s", *output.Errors[0].Message)
		}
	}
	for _, delObj := range output.Deleted {
		err = s3.NewObjectNotExistsWaiter(time_web_s3.S3Client).Wait(
			ctx, &s3.HeadObjectInput{Bucket: aws.String(configs.AppConfig.S3.Bucket), Key: delObj.Key}, time.Minute)
		if err != nil {
			log.Printf("Ошибка ожидания удаления объекта %s.\n", *delObj.Key)
		} else {
			log.Printf("Удалён %s.\n", *delObj.Key)
		}
	}
	return nil
}

func GetFileURL(objectKeys string) string {
	encodedKey := url.PathEscape(objectKeys)
	return fmt.Sprintf("%s/%s/%s", configs.AppConfig.S3.Endpoint, configs.AppConfig.S3.Bucket, encodedKey)
}

func ListFiles(prefix string) (*s3.ListObjectsV2Output, error) {
	res, err := time_web_s3.S3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(configs.AppConfig.S3.Bucket),
		Prefix: aws.String(prefix),
	})
	if err != nil {
		return nil, fmt.Errorf("ошибка получения списка файлов: %w", err)
	}
	return res, nil
}

func DownloadFile(ctx context.Context, objectKey string) ([]byte, error) {
	objectKey = strings.TrimLeft(objectKey, "/") // Убираем ведущий слеш
	result, err := time_web_s3.S3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(configs.AppConfig.S3.Bucket),
		Key:    aws.String(objectKey), // Используем чистый objectKey
	})
	if err != nil {
		var noKeyErr *types.NoSuchKey
		if errors.As(err, &noKeyErr) {
			return nil, fmt.Errorf("файл с ключом %s не найден: %w", objectKey, err)
		}
		return nil, fmt.Errorf("ошибка получения файла: %w", err)
	}
	defer result.Body.Close()

	data, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}
	return data, nil
}

func ListBuckets(ctx context.Context) ([]types.Bucket, error) {
	var err error
	var output *s3.ListBucketsOutput
	var buckets []types.Bucket
	bucketPaginator := s3.NewListBucketsPaginator(time_web_s3.S3Client, &s3.ListBucketsInput{})
	for bucketPaginator.HasMorePages() {
		output, err = bucketPaginator.NextPage(ctx)
		if err != nil {
			var apiErr smithy.APIError
			if errors.As(err, &apiErr) && apiErr.ErrorCode() == "AccessDenied" {
				fmt.Println("You don't have permission to list buckets for this account.")
				err = apiErr
			} else {
				log.Printf("Couldn't list buckets for your account. Here's why: %v\n", err)
			}
			break
		} else {
			buckets = append(buckets, output.Buckets...)
		}
	}
	return buckets, err
}
