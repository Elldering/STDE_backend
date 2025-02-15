package main

import (
	"STDE_proj/configs"
	"STDE_proj/internal/routes"
	"STDE_proj/utils/db"
	"STDE_proj/utils/time_web_s3"
	"github.com/gin-gonic/gin"
	"log"
)

// Запускает приложение. Точка входа
func main() {
	// Чтение и загрузка конфигурации приложение - загрузка config.yaml
	if err := configs.LoadConfig("env"); err != nil {
		log.Fatalf("Не удалось загрузить конфигурацию: %v", err)
	}

	// Подключение к базе данных
	if err := db.Connect(); err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	if err := time_web_s3.InitS3(); err != nil {
		log.Fatalf("Не удалось подключиться к S3: %v", err)
	}

	// Инициализация экземпляра маршрутизатора
	router := gin.Default()

	log.Println(configs.AppConfig.YandexDisk.Token)

	// Использование конфигурации CORS в нашем маршрутизаторе
	router.Use(configs.CorsConfig())

	// Передаём маршруты для дальнейшей обработки их маршрутизатором
	routes.Routes(router)

	// Запускаем наш маршрутизатор (приложение)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
