package tests

import (
	"STDE_proj/configs"
	"testing"
)

// TestLoadConfig тестирование выгрузки данных с конфига используя функции LoadConfig и структуры Config
// (AppConfig глобальная переменная, которая ссылается на структуру Config)
func TestLoadConfig(t *testing.T) {
	err := configs.LoadConfig("../configs/config.yaml") // Указываем путь к файлу конфигурации
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
