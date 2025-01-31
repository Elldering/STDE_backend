package tests

import (
	"STDE_proj/utils/db"
	"testing"
)

func TestConnectDB(t *testing.T) {
	err := db.Connect()
	if err != nil {
		t.Fatalf("Ошибка подключение к базе данных: %v", err)
	}
}
