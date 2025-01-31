package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"STDE_proj/configs"
	"STDE_proj/internal/routes"
	"STDE_proj/utils/db"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// SetupTestRouter настраивает тестовый маршрутизатор Gin
func SetupTestRouter() *gin.Engine {
	// Проверка текущей рабочей директории
	dir, err := os.Getwd()
	if err != nil {
		panic("Не удалось получить текущую рабочую директорию: " + err.Error())
	}
	println("Текущая рабочая директория:", dir)

	// Загружаем конфигурацию
	if err := configs.LoadConfig("../../../configs/config.yaml"); err != nil {
		panic("Не удалось загрузить конфигурацию: " + err.Error())
	}

	// Подключаемся к базе данных
	if err := db.Connect(); err != nil {
		panic("Ошибка подключения к базе данных: " + err.Error())
	}

	// Создаем маршрутизатор
	router := gin.Default()

	// Настраиваем маршруты
	routes.Routes(router)

	return router
}

// TestGetUserGroups тестирует ручку GET /usergroups
func TestGetUserGroups(t *testing.T) {
	router := SetupTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/usergroups", nil)
	router.ServeHTTP(w, req)

	// Проверка кода состояния HTTP
	assert.Equal(t, http.StatusOK, w.Code)

	// Проверка структуры данных
	var responseBody []map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	require.NoError(t, err)

	// Проверка, что ответ содержит массив объектов
	assert.IsType(t, []map[string]interface{}{}, responseBody)

	// Проверка первого элемента массива (если есть данные)
	if len(responseBody) > 0 {
		assert.Contains(t, responseBody[0], "id")
		assert.Contains(t, responseBody[0], "name")
	}
}
