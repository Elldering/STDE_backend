package test_config

import (
	"STDE_proj/tests"
	"testing"
	"time"

	"STDE_proj/configs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestLoadConfig тестирование выгрузки данных с конфига используя функции LoadConfig и структуры Config
func TestLoadConfig(t *testing.T) {
	err := configs.LoadConfig("env.test") // Указываем путь к файлу конфигурации
	if err != nil {
		t.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	if configs.AppConfig.Database.Host == "" {
		t.Errorf("DBHost ничего в себе не содержит или отсутствует")
	}

	if configs.AppConfig.Database.Port == 0 {
		t.Errorf("DBPort ничего в себе не содержит или отсутствует")
	}

	if configs.AppConfig.Database.User == "" {
		t.Errorf("DBUser ничего в себе не содержит или отсутствует")
	}

	if configs.AppConfig.Database.Password == "" {
		t.Errorf("DBPass ничего в себе не содержит или отсутствует")
	}

	if configs.AppConfig.Database.Name == "" {
		t.Errorf("DBName ничего в себе не содержит или отсутствует")
	}

}

// TestLoadCorsConfig тестирует загрузку конфигурации CORS
func TestLoadCorsConfig(t *testing.T) {
	// Загрузить конфигурацию
	err := configs.LoadConfig("env.test")
	require.NoError(t, err, "Ошибка загрузки конфигурации")

	// Проверка значений полей конфигурации CORS
	corsConfig := configs.AppConfig.CORS

	// Проверка Origins
	assert.NotEmpty(t, corsConfig.AllowOrigins, "AllowOrigins не должен быть пустым")
	for _, origin := range tests.AllowedOriginsTest {
		assert.Contains(t, corsConfig.AllowOrigins, origin, "AllowOrigins должен содержать '"+origin+"'")
	}

	// Проверка Methods
	assert.NotEmpty(t, corsConfig.AllowMethods, "AllowMethods не должен быть пустым")
	for _, method := range tests.AllowedMethodsTest {
		assert.Contains(t, corsConfig.AllowMethods, method, "AllowMethods должен содержать '"+method+"'")
	}

	// Проверка Headers
	assert.NotEmpty(t, corsConfig.AllowHeaders, "AllowHeaders не должен быть пустым")
	for _, header := range tests.AllowedHeadersTest {
		assert.Contains(t, corsConfig.AllowHeaders, header, "AllowHeaders должен содержать '"+header+"'")
	}

	// Проверка Credentials
	assert.Equal(t, true, corsConfig.AllowCredentials, "AllowCredentials должен быть true")

	// Проверка MaxAge
	expectedMaxAge := 12 * time.Hour / time.Second // 12 часов в секундах
	assert.Equal(t, expectedMaxAge, time.Duration(corsConfig.MaxAge), "MaxAge должен быть 43200 секунд")
}
