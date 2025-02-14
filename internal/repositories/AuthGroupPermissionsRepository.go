package repositories

import (
	"STDE_proj/internal/models"
	"STDE_proj/utils/db"
	"fmt"
	"log"
)

// Функция для создания связи группы и прав доступа
func PostAuthGroupPermission(agp models.AuthGroupPermissions) error {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return fmt.Errorf("подключение к базе данных не инициализировано")
	}

	query := "INSERT INTO auth_group_permissions (group_id, permission_id) VALUES ($1, $2)"
	_, err := db.DB.Exec(query, agp.GroupID, agp.PermissionID)
	return err
}

//
//func DeleteAuthGroupPermission(agp models.AuthGroupPermissions) error {
//	if db.DB == nil {
//
//	}
//}
//
//func UpdateAuthGroupPermission(agp models.AuthGroupPermissions) error {
//	if db.DB == nil {
//	}
//}
//
//func GetAuthGroupPermission(agp models.AuthGroupPermissions) (models.AuthGroupPermissions, error) {
//
//}
