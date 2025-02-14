package services

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories"
)

func GetAllUserProfile() ([]models.UserProfile, error) {
	return repositories.GetAllUserProfile()
}

func GetUserProfileById(id int) (models.UserProfile, error) {
	return repositories.GetUserProfileById(id)
}

func PostUserProfile(agp models.UserProfile) error {
	return repositories.PostUserProfile(agp)
}

func PutUserProfile(id int, agp models.UserProfile) error {
	return repositories.PutUserProfile(id, agp)
}

func DeleteUserProfile(id int) error {
	return repositories.DeleteUserProfile(id)
}
