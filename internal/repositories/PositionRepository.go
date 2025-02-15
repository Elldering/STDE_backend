package repositories

import (
	"STDE_proj/internal/models"
	"STDE_proj/utils/db"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
)

func GetAllPositions() ([]models.Position, error) {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return nil, fmt.Errorf("подключение к базе данных не инициализировано")
	}

	rows, err := db.DB.Query("SELECT id, name, description, price, available, image FROM position")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var positions []models.Position
	for rows.Next() {
		var position models.Position
		if err := rows.Scan(&position.ID, &position.Name, &position.Description, &position.Price, &position.Available, &position.Image); err != nil {
			return nil, err
		}
		positions = append(positions, position)
	}
	return positions, nil
}

func GetPositionById(id int) (models.Position, error) {
	row := db.DB.QueryRow("SELECT id, name, description, price, available, image FROM position WHERE id=$1", id)

	var position models.Position
	if err := row.Scan(&position.ID, &position.Name, &position.Description, &position.Price, &position.Available, &position.Image); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Position{}, fmt.Errorf("блюдо с id %d не найдена", id)
		}
		return models.Position{}, err
	}
	return position, nil
}

// Создание новой позиции
func PostPosition(agp models.Position) error {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return errors.New("подключение к базе данных не инициализировано")
	}

	if strings.TrimSpace(agp.Description) == "" || strings.TrimSpace(agp.Name) == "" || strings.TrimSpace(agp.Image) == "" {
		log.Println("Ошибка: пустые поля codename или description")
		return errors.New("поля Name, Image и Description не могут быть пустыми")
	}
	query := "INSERT INTO position (name, description, price, available, image) VALUES ($1, $2, $3, $4, $5)"
	result, err := db.DB.Exec(query, agp.Name, agp.Description, agp.Price, agp.Available, agp.Image)

	if err != nil {
		log.Printf("Ошибка при добавлении блюда: %v", err)
		return fmt.Errorf("ошибка при добавлении блюда: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Ошибка при проверке количества добавленных строк: %v", err)
		return fmt.Errorf("ошибка при проверке количества добавленных строк: %v", err)
	}

	if rowsAffected == 0 {
		log.Println("Ошибка: блюдо не было добавлено")
		return errors.New("блюдо не было добавлено")
	}

	return err
}

func PutPosition(id int, position models.Position) error {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return errors.New("подключение к базе данных не инициализировано")
	}
	query := "UPDATE position SET name=$1, description=$2, price=$3, available=$4, image=$5 WHERE id=$6"
	result, err := db.DB.Exec(query, position.Name, position.Description, position.Price, position.Available, position.Image, id)
	if err != nil {
		log.Printf("Ошибка при удалении блюда с id %d: %v", id, err)
		return fmt.Errorf("ошибка при удалении блюда: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Ошибка при проверке количества удаленных строк для id %d: %v", id, err)
		return fmt.Errorf("ошибка при проверке количества удаленных строк: %v", err)
	}
	if rowsAffected == 0 {
		log.Printf("Ошибка: блюдо с id %d не найдено", id)
		return fmt.Errorf("блюдо с id %d не найдено", id)
	}

	return nil

}

func DeletePosition(id int) error {
	_, err := db.DB.Exec("DELETE FROM position WHERE id=$1", id)
	return err
}
