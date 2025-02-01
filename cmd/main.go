package main

import (
	"STDE_proj/configs"
	"STDE_proj/internal/routes"
	"STDE_proj/utils/db"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	if err := configs.LoadConfig("configs/config.yaml"); err != nil {
		log.Fatalf("Не удалось загрузить конфигурацию: %v", err)
	}

	if err := db.Connect(); err != nil {
		log.Fatalf("Не удалось подключиться у базе данных: %v", err)
	}

	router := gin.Default()

	router.Use(configs.CorsConfig())

	routes.Routes(router)

	router.Run(":8080")
}
