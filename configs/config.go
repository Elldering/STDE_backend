package configs

import (
	"fmt"
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

var AppConfig *Config

// LoadConfig загружает конфигурационные данные и инициализирует AppConfig
func LoadConfig(path string) error {
	viper.SetConfigFile(path)
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

// CorsConfig возвращает конфигурацию CORS
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
