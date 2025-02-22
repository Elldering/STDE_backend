package RegisterRepository

import (
	"STDE_proj/utils/database"
	"log"
)

func GetVerificationCode(userID int, code int) (int, error) {
	query := "SELECT code FROM verify_code WHERE auth_user_id = $1"
	err := database.DB.QueryRow(query, userID).Scan(&code)
	if err != nil {
		log.Printf("Ошибка при получении кода подтверждения: %v", err)
		return code, err
	}
	return code, nil
}

func UpdateEmailVerified(userID int) error {
	query := "UPDATE auth_user SET is_email_verify = TRUE WHERE id = $1"
	_, err := database.DB.Exec(query, userID)
	if err != nil {
		log.Printf("Ошибка при обновлении статуса подтверждения почты: %v", err)
		return err
	}
	return nil
}

func DeleteVerificationCode(userID int) error {
	query := "DELETE FROM verify_code WHERE auth_user_id = $1"
	_, err := database.DB.Exec(query, userID)
	if err != nil {
		log.Printf("Ошибка при удалении кода: %v", err)
		return err
	}
	return nil
}
