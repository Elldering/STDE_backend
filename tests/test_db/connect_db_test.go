package test_db

import (
	"testing"

	"STDE_proj/configs"
	"STDE_proj/utils/db"
	"github.com/stretchr/testify/assert"
)

// TestConnectDB тестирует подключение к базе данных
func TestConnectDB(t *testing.T) {
	// Загружаем конфигурацию
	err := configs.LoadConfig("../../configs/config.yaml")
	assert.NoError(t, err, "Ошибка загрузки конфигурации")

	// Подключаемся к базе данных
	err = db.Connect()
	assert.NoError(t, err, "Ошибка подключения к базе данных")
}
