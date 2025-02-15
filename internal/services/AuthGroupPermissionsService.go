package services

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories"
)

func GetAuthGroupPermissions() ([]models.AuthGroupPermissions, error) {
	return repositories.GetAuthGroupPermissions()
}

func PostAuthGroupPermission(agp models.AuthGroupPermissions) error {
	return repositories.PostAuthGroupPermission(agp)
}

func GetAuthGroupPermissionsId(id int) (models.AuthGroupPermissions, error) {
	return repositories.GetAuthGroupPermissionsId(id)
}

func PutAuthGroupPermissions(id int, data models.AuthGroupPermissions) error {
	return repositories.PutAuthGroupPermissions(id, data)
}

func DeleteAuthGroupPermissions(id int) error {
	return repositories.DeleteAuthGroupPermissions(id)
}
