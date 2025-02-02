package services

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories"
)

func CreateAuthGroupPermission(agp models.AuthGroupPermissions) error {
	return repositories.CreateAuthGroupPermission(agp)
}
