package test_db

import (
	"STDE_proj/utils/database"
	"testing"

	"STDE_proj/configs"
	"github.com/stretchr/testify/assert"
)

// TestConnectDB тестирует подключение к базе данных
func TestConnectDB(t *testing.T) {
	// Загружаем конфигурацию
	err := configs.LoadConfig("env.test")
	assert.NoError(t, err, "Ошибка загрузки конфигурации")

	// Подключаемся к базе данных
	err = database.Connect()
	assert.NoError(t, err, "Ошибка подключения к базе данных")
}
