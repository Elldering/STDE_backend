package services

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories"
)

func GetOrderPositions() ([]models.OrderPosition, error) {
	return repositories.GetOrderPositions()
}

func GetOrderPositionById(id int64) (models.OrderPosition, error) {
	return repositories.GetOrderPositionById(id)
}

func PostOrderPosition(orderPosition models.OrderPosition) error {
	return repositories.PostOrderPosition(orderPosition)
}

func PutOrderPosition(id int64, orderPosition models.OrderPosition) error {
	return repositories.PutOrderPosition(id, orderPosition)
}

func DeleteOrderPosition(id int64) error {
	return repositories.DeleteOrderPosition(id)
}
