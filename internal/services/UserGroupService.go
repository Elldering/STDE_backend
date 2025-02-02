package services

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories"
)

func GetAllUserGroups() ([]models.UserGroup, error) {
	return repositories.GetAllUserGroups()
}

func GetUserGroupById(id int) (models.UserGroup, error) {
	return repositories.GetUserGroupById(id)
}

func PostUserGroup(agp models.UserGroup) error {
	return repositories.PostUserGroup(agp)
}
func PutUserGroup(id int, agp models.UserGroup) error {
	return repositories.PutUserGroup(id, agp)
}
func DeleteUserGroup(id int) error {
	return repositories.DeleteUserGroup(id)
}
