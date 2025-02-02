package services

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories"
)

func PostAuthGroupPermission(agp models.AuthGroupPermissions) error {
	return repositories.CreateAuthGroupPermission(agp)
}
