package services

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories"
)

func GetDocumentAuthUsers() ([]models.DocumentAuthUser, error) {
	return repositories.GetDocumentAuthUsers()
}

func GetDocumentAuthUserById(id int64) (models.DocumentAuthUser, error) {
	return repositories.GetDocumentAuthUserById(id)
}

func PostDocumentAuthUser(docAuthUser models.DocumentAuthUser) error {
	return repositories.PostDocumentAuthUser(docAuthUser)
}

func PutDocumentAuthUser(id int64, docAuthUser models.DocumentAuthUser) error {
	return repositories.PutDocumentAuthUser(id, docAuthUser)
}

func DeleteDocumentAuthUser(id int64) error {
	return repositories.DeleteDocumentAuthUser(id)
}
