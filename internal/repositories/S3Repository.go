package repositories

import (
	"STDE_proj/configs"
	"STDE_proj/utils/time_web_s3"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
	"mime/multipart"
	"net/url"
)

func PostFile(file multipart.File, fileName string) (string, error) {
	_, err := time_web_s3.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(configs.AppConfig.S3.Bucket),
		Key:    aws.String(fileName),
		Body:   file,
	})
	if err != nil {
		return "", fmt.Errorf("ошибка загрузки файла: %w", err)
	}
	encodedFileName := url.PathEscape(fileName)
	fileURL := fmt.Sprintf("%s/%s/%s", configs.AppConfig.S3.Endpoint, configs.AppConfig.S3.Bucket, encodedFileName)
	return fileURL, nil
}

// DeleteFile удаляет файл из S3.
func DeleteFile(fileName string) error {
	_, err := time_web_s3.S3Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(configs.AppConfig.S3.Bucket),
		Key:    aws.String(fileName),
	})
	if err != nil {
		return fmt.Errorf("ошибка удаления файла: %w", err)
	}
	return nil
}

// GetFileURL генерирует URL для доступа к файлу.
func GetFileURL(fileName string) string {
	return fmt.Sprintf("%s/%s/%s", configs.AppConfig.S3.Endpoint, configs.AppConfig.S3.Bucket, fileName)
}

// ListFiles возвращает список файлов в бакете с указанным префиксом.
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

// DownloadFile скачивает файл из S3 и возвращает его содержимое в виде среза байт.
func DownloadFile(fileName string) ([]byte, error) {
	result, err := time_web_s3.S3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(configs.AppConfig.S3.Bucket),
		Key:    aws.String(fileName),
	})
	if err != nil {
		return nil, fmt.Errorf("ошибка получения файла: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(result.Body)

	data, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}
	return data, nil
}

// ListBuckets возвращает список бакетов, доступных в аккаунте.
func ListBuckets() (*s3.ListBucketsOutput, error) {
	res, err := time_web_s3.S3Client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		return nil, fmt.Errorf("ошибка получения списка бакетов: %w", err)
	}
	return res, nil
}
