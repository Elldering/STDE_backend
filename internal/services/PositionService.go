package services

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories"
)

func GetAllPositions() ([]models.Position, error) {
	return repositories.GetAllPositions()
}

func GetPositionById(id int) (models.Position, error) {
	return repositories.GetPositionById(id)
}

func PostPosition(agm models.Position) error {
	return repositories.PostPosition(agm)
}

func PutPosition(id int, agm models.Position) error {
	return repositories.PutPosition(id, agm)
}

func DeletePosition(id int) error {
	return repositories.DeletePosition(id)
}
