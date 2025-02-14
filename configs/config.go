package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Config структура для хранения всех конфигурационных данных
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Logging  LoggingConfig  `mapstructure:"logging"`
	CORS     CORSConfig     `mapstructure:"cors"`
	S3       S3Config       `mapstructure:"s3"`
}

// ServerConfig структура для конфигурации сервера
type ServerConfig struct {
	GRPCPort int `mapstructure:"grpc_port"`
}

// DatabaseConfig структура для конфигурации базы данных
type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

// LoggingConfig структура для конфигурации логирования
type LoggingConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

// CORSConfig структура для конфигурации CORS
type CORSConfig struct {
	AllowOrigins     []string `mapstructure:"allow_origins"`
	AllowMethods     []string `mapstructure:"allow_methods"`
	AllowHeaders     []string `mapstructure:"allow_headers"`
	AllowCredentials bool     `mapstructure:"allow_credentials"`
	MaxAge           int      `mapstructure:"max_age"`
}

// S3Config структура для конфигурации S3
type S3Config struct {
	Bucket    string `mapstructure:"bucket"`
	Region    string `mapstructure:"region"`
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
	Endpoint  string `mapstructure:"endpoint"`
}

var AppConfig *Config

// LoadConfig загружает конфигурационные данные и инициализирует AppConfig.
// Принимает параметр envFilePath, который указывает тип файла .env:
// "env" для production конфигурации и "env.test" для тестовой конфигурации.
// Возвращает ошибку, если произошли проблемы с загрузкой файла .env или файла конфигурации.
func LoadConfig(envFilePath string) error {

	var envPathLoad string

	// Обработка выбора между test и prod .env
	switch envFilePath {
	case "env":
		envPathLoad = ".env"
	case "env.test":
		envPathLoad = "../../.env.test"
	default:
		log.Println("Вы ввели неправильный тип файла (env) или он отсутствует в системе")
	}

	// Загрузка переменных окружения из указанного файла .env
	err := godotenv.Load(envPathLoad)
	if err != nil {
		return fmt.Errorf("ошибка при загрузке файла .env, %w", err)
	}

	// Получаем путь к конфигурационному файлу из переменной окружения
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		return fmt.Errorf("переменная окружения CONFIG_PATH не установлена")
	}

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	// Читаем файл конфигурации
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("ошибка при чтении конфигурационного файла, %w", err)
	}

	// Инициализируем глобальную переменную AppConfig
	AppConfig = &Config{}

	// Распаковываем значения в структуру
	if err := viper.Unmarshal(&AppConfig); err != nil {
		return fmt.Errorf("ошибка при распаковке данных в структуру, %w", err)
	}

	return nil
}

// CorsConfig возвращает middleware для настройки CORS (Cross-Origin Resource Sharing) в Gin.
// Функция использует настройки CORS из глобальной переменной AppConfig.
// Если список AllowOrigins пуст, функция вызывает panic.
//
// Возвращаемое значение:
//   - gin.HandlerFunc: middleware для настройки CORS.
func CorsConfig() gin.HandlerFunc {
	config := cors.Config{
		AllowOrigins:     AppConfig.CORS.AllowOrigins,
		AllowMethods:     AppConfig.CORS.AllowMethods,
		AllowHeaders:     AppConfig.CORS.AllowHeaders,
		AllowCredentials: AppConfig.CORS.AllowCredentials,
		MaxAge:           time.Duration(AppConfig.CORS.MaxAge) * time.Second,
	}

	// Проверяем, что AllowOrigins не пустой
	if len(config.AllowOrigins) == 0 {
		panic("CORS настройка: все origins отключены")
	}

	return cors.New(config)
}
