package main

import (
	"STDE_proj/configs"
	"STDE_proj/internal/routes"
	"STDE_proj/utils/db"
	"STDE_proj/utils/time_web_s3"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"log"
	"os"
)

// Запускает приложение. Точка входа
func main() {
	// Чтение и загрузка конфигурации приложение - загрузка config.yaml
	if err := configs.LoadConfig("env"); err != nil {
		log.Fatalf("Не удалось загрузить конфигурацию: %v", err)
	}
	IsTestRoutes := os.Getenv("TEST_ROUTES")
	// Подключение к базе данных
	if err := db.Connect(); err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	if err := time_web_s3.InitS3(); err != nil {
		log.Fatalf("Не удалось подключиться к S3: %v", err)
	}

	cron := cron.New()
	cron.AddFunc("@every 15m", func() {
		err := db.DeleteNoRegUser()
		if err != nil {
			log.Printf("Ошибка при вызове функции: %s", err)
		}
	})

	// Инициализация экземпляра маршрутизатора
	router := gin.Default()

	// Использование конфигурации CORS в нашем маршрутизаторе
	router.Use(configs.CorsConfig())

	// Передаём маршруты для дальнейшей обработки их маршрутизатором
	routes.Routes(router)
	if IsTestRoutes == "true" {
		routes.TestRoutes(router)
		log.Println("Конфигурация тестовых маршрутов включена")
	} else {
		log.Println("Конфигурация тестовых маршрутов отключена")
	}
	cron.Start()
	// Запускаем наш маршрутизатор (приложение)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
