package time_web_s3

import (
	"STDE_proj/configs"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// S3Client TODO переписать. Сделать вместо глобальной переменной, подключение СУКАААА!!!!
var S3Client *s3.Client

// InitS3 инициализирует S3-клиент, используя aws-sdk-go-v2.
func InitS3() error {
	var err error
	// Загружаем конфигурацию
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(configs.AppConfig.S3.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			configs.AppConfig.S3.AccessKey,
			configs.AppConfig.S3.SecretKey,
			"",
		)),
		config.WithEndpointResolver(aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL:           configs.AppConfig.S3.Endpoint,
				SigningRegion: configs.AppConfig.S3.Region,
			}, nil
		})),
	)
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации S3: %v", err)
	}

	// Важно: используем Path-style URL, так как Virtual-hosted-style может не работать в Timeweb
	S3Client = s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})
	return nil
}
