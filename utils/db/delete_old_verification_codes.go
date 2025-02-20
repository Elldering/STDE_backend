package db

import "log"

func DeleteNoRegUser() error {
	query := "SELECT delete_inactive_users();"
	_, err := DB.Exec(query)
	if err != nil {
		log.Printf("Ошибка при удалении пользователей, которые не подтвердили почту: %v", err)
		return err
	}
	log.Println("Пользователь, который не подтвердил почту успешно удален")
	return nil
}
