package crons

import (
	"STDE_proj/utils/database"
	"github.com/robfig/cron/v3"
	"log"
)

// Crones хранит в себе все исполняемые функции
// которые в последствии мы вызываем в main.go
func Crones(cron *cron.Cron) error {
	_, err := cron.AddFunc("@every 11m", func() {
		err := database.DeleteNoRegUser()
		if err != nil {
			log.Printf("Ошибка при вызове функции: %s", err)
		}
	})
	if err != nil {
		return err
	}

	_, err = cron.AddFunc("@every 2d", func() {
		err := database.DeleteOldInvalidTokens()
		if err != nil {
			log.Printf("Ошибка при вызове функции: %s", err)
		}
	})
	return nil
}
