package db

import "log"

func GetVerificationCode(userID int) (string, error) {
	var code string
	query := "SELECT code, auth_user_id FROM verify_code WHERE auth_user_id = $1"
	err := DB.QueryRow(query, userID).Scan(&code)
	if err != nil {
		log.Printf("Ошибка при получении кода подтверждения: %v", err)
		return "", err
	}
	return code, nil
}

// DeleteExpiredVerificationCodes удаляет просроченные коды подтверждения
func DeleteExpiredVerificationCodes() error {
	query := "DELETE FROM verify_code WHERE created_at < NOW() - INTERVAL '5 seconds'"
	_, err := DB.Exec(query)
	if err != nil {
		log.Printf("Ошибка при удалении просроченных кодов: %v", err)
		return err
	}
	log.Println("Просроченные коды подтверждения успешно удалены")
	return nil
}

func DeleteOldUser() error {
	query := "DELETE FROM auth_user WHERE "
	_, err := DB.Exec(query)
	if err != nil {
		log.Printf("Ошибка при удалении просроченных кодов: %v", err)
		return err
	}
	log.Println("Просроченные коды подтверждения успешно удалены")
	return nil
}
