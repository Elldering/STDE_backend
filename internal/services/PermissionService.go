package services

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories"
)

func GetPermissions() ([]models.Permission, error) {
	return repositories.GetAllPermissions()
}

func GetPermissionById(id int) (models.Permission, error) {
	return repositories.GetPermissionById(id)
}

func PostPermission(agp models.Permission) error {
	return repositories.PostPermission(agp)
}

func PutPermission(id int, agp models.Permission) error {
	return repositories.PutPermission(id, agp)
}

func DeletePermission(id int) error {
	return repositories.DeletePermission(id)
}
