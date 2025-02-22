package services

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories"
)

func GetMenuPositions() ([]models.MenuPosition, error) {
	return repositories.GetMenuPositions()
}

func GetMenuPositionById(id int64) (models.MenuPosition, error) {
	return repositories.GetMenuPositionById(id)
}

func PostMenuPosition(menuPosition models.MenuPosition) error {
	return repositories.PostMenuPosition(menuPosition)
}

func PutMenuPosition(id int64, menuPosition models.MenuPosition) error {
	return repositories.PutMenuPosition(id, menuPosition)
}

func DeleteMenuPosition(id int64) error {
	return repositories.DeleteMenuPosition(id)
}
