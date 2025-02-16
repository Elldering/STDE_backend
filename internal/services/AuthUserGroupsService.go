package services

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories"
)

func GetAuthUserGroupsAll() ([]models.AuthUserGroups, error) {
	return repositories.GetAuthUserGroupsAll()
}

func GetAuthUserGroupsById(id int) (models.AuthUserGroups, error) {
	return repositories.GetAuthUserGroupsById(id)
}

func PostAuthUserGroups(data models.AuthUserGroups) error {
	return repositories.PostAuthUserGroups(data)
}

func PutAuthUserGroups(id int, data models.AuthUserGroups) error {
	return repositories.PutAuthUserGroups(id, data)
}

func DeleteAuthUserGroups(id int) error {
	return repositories.DeleteAuthUserGroups(id)
}
