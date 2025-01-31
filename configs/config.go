package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

// Config структура для хранения всех конфигурационных данных
type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
}

// DatabaseConfig структура для конфигурации базы данных
type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

// AppConfig глобальная переменная, которая ссылается на структуру Config
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
