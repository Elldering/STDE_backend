package database

import (
	"log"
)

func DeleteNoRegUser() error {
	// Исправленный SQL-запрос
	query := `
        DELETE FROM auth_user
        WHERE is_email_verify = FALSE
          AND id IN (
              SELECT auth_user_id
              FROM verify_code
              WHERE created_at < NOW() - INTERVAL '10 minutes'
          );`

	// Выполняем запрос
	result, err := DB.Exec(query)
	if err != nil {
		log.Printf("Ошибка при удалении пользователей, которые не подтвердили почту: %v", err)
		return err
	}

	// Получаем количество удалённых строк
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Ошибка при получении количества удалённых строк: %v", err)
		return err
	}

	// Логируем результат
	if rowsAffected > 0 {
		log.Printf("Удалено %d пользователей, которые не подтвердили почту.", rowsAffected)
	} else {
		log.Println("Пользователей для удаления не найдено.")
	}

	return nil
}

func DeleteOldInvalidTokens() error {
	// SQL-запрос для удаления токенов старше 5 дней
	query := `DELETE FROM invalid_tokens WHERE created_at < NOW() - INTERVAL '5 days';`

	// Выполняем запрос
	result, err := DB.Exec(query)
	if err != nil {
		log.Printf("Ошибка при удалении просроченных токенов: %v", err)
		return err
	}

	// Получаем количество удалённых строк
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Ошибка при получении количества удалённых строк: %v", err)
		return err
	}

	// Логируем результат
	if rowsAffected > 0 {
		log.Printf("Удалено %d просроченных токенов.", rowsAffected)
	} else {
		log.Println("Просроченных токенов для удаления не найдено.")
	}

	return nil
}
