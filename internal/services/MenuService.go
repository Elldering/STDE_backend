package services

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories"
)

func GetMenu() ([]models.Menu, error) {
	return repositories.GetMenu()
}
func GetMenuById(id int) (models.Menu, error) {
	return repositories.GetMenuById(id)
}
func PostMenu(agm models.Menu) error {
	return repositories.PostMenu(agm)
}
func PutMenu(id int, agm models.Menu) error {
	return repositories.PutMenu(id, agm)
}
func DeleteMenu(id int) error {
	return repositories.DeleteMenu(id)
}
