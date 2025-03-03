package repositories

import (
	"STDE_proj/internal/models"
	"STDE_proj/utils/database"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/gin-gonic/gin"
	"log"
)

func GetAllUserProfile() ([]models.UserProfile, error) {

	rows, err := database.DB.Query("SELECT id, first_name, last_name, middle_name, rating, profile_image FROM user_profile")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userProfiles []models.UserProfile
	for rows.Next() {
		var userProfile models.UserProfile
		if err := rows.Scan(&userProfile.ID, &userProfile.FirstName, &userProfile.LastName, &userProfile.MiddleName, &userProfile.Rating, &userProfile.ProfileImage); err != nil {
			return nil, err
		}
		userProfiles = append(userProfiles, userProfile)
	}

	return userProfiles, nil
}

func GetUserProfileById(id int) (models.UserProfile, error) {

	row := database.DB.QueryRow("SELECT id, first_name, last_name, middle_name, rating, profile_image FROM user_profile WHERE id=$1", id)
	var userProfile models.UserProfile
	if err := row.Scan(&userProfile.ID, &userProfile.FirstName, &userProfile.LastName, &userProfile.MiddleName, &userProfile.Rating, &userProfile.ProfileImage); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.UserProfile{}, fmt.Errorf("группа профиля пользователя с id %d не найдена", id)
		}
		return models.UserProfile{}, err
	}
	return userProfile, nil
}

func PostUserProfile(agp models.UserProfile) error {

	query := "INSERT INTO user_profile (first_name, last_name, middle_name, rating, profile_image) VALUES ($1, $2, $3, $4, $5)"
	result, err := database.DB.Exec(query, agp.FirstName, agp.LastName, agp.MiddleName, agp.Rating, agp.ProfileImage)
	if err != nil {
		log.Printf("Ошибка при добавлении UserProfile: %v", err)
		return fmt.Errorf("ошибка при добавлении UserProfile: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Ошибка при проверке количества добавленных строк: %v", err)
		return fmt.Errorf("ошибка при проверке количества добавленных строк: %v", err)
	}
	if rowsAffected == 0 {
		log.Println("Ошибка: UserProfile не была добавлена")
		return errors.New("UserProfile не была добавлена")
	}

	return nil
}

func PutUserProfile(id int, agp models.UserProfile) error {

	row := database.DB.QueryRow("SELECT id, first_name, last_name, middle_name, rating, profile_image FROM user_profile WHERE id=$1", id)
	var userGroup models.UserProfile
	if err := row.Scan(&userGroup.ID, &userGroup.FirstName, &userGroup.LastName, &userGroup.MiddleName, &userGroup.Rating, &userGroup.ProfileImage); err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("профиль пользователя с id %d не найдена", id)
		}
		return err
	}

	result, err := database.DB.Exec("UPDATE user_profile SET first_name=$2, last_name=$3, middle_name=$4, rating=$5, profile_image=$6 WHERE id=$1", id, agp.FirstName, agp.LastName, agp.MiddleName, agp.Rating, agp.ProfileImage)
	if err != nil {
		log.Println("Ошибка при обновлении профилей пользователей:", err)
		return fmt.Errorf("ошибка при обновлении профилей пользователей: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Ошибка при проверке количества обновленных строк:", err)
		return fmt.Errorf("ошибка при проверке количества обновленных строк: %v", err)
	}
	if rowsAffected == 0 {

		return fmt.Errorf("профиль пользователя с id %d не найдена", id)
	}

	return nil
}

func DeleteUserProfile(id int) error {

	result, err := database.DB.Exec("DELETE FROM user_profile WHERE id=$1", id)
	if err != nil {
		log.Println("Ошибка при удалении профиля пользователя:", err)
		return fmt.Errorf("ошибка при удалении профиля пользователя: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Ошибка при проверке количества удаленных строк:", err)
		return fmt.Errorf("ошибка при проверке количества удаленных строк: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("группа пользователей с id %d не найдена", id)
	}

	return nil
}
