package services

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories"
)

func GetUserDocuments() ([]models.UserDocument, error) {
	return repositories.GetUserDocuments()
}

func GetUserDocumentById(id int) (models.UserDocument, error) {
	return repositories.GetUserDocumentById(id)
}

func PostUserDocument(userDoc models.UserDocument) error {
	return repositories.PostUserDocument(userDoc)
}

func PutUserDocument(id int, userDoc models.UserDocument) error {
	return repositories.PutUserDocument(id, userDoc)
}

func DeleteUserDocument(id int) error {
	return repositories.DeleteUserDocument(id)
}
