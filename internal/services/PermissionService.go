package services

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories"
)

func GetPermissions() ([]models.Permisson, error) {
	return repositories.GetAllPermissions()
}

func GetPermissionById(id int) (models.Permisson, error) {
	return repositories.GetPermissionById(id)
}

func PostPermission(agp models.Permisson) error { return repositories.PostPermission(agp) }

func PutPermission(id int, agp models.Permisson) error { return repositories.PutPermission(id, agp) }

func DeletePermission(id int) error {
	return repositories.DeletePermission(id)
}
