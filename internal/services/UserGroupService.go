package services

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories"
)

func GetAllUserGroups() ([]models.UserGroup, error) {
	return repositories.GetAllUserGroups()
}
