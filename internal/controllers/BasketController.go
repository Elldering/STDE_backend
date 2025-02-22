package controllers

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func GetBasketsHandler(c *gin.Context) {
	data, err := services.GetBasket()
	if err != nil {
		log.Printf("Ошибка получения корзины: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Ошибка получения корзины"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetBasketByIdUserHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Неверный формат ID: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}
	data, err := services.GetBasketById(id)
	if err != nil {
		log.Printf("Ошибка при получении корзны пользователя с id :%v, %v", id, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при получении корзины пользователя"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func PostBasketHandler(c *gin.Context) {
	var agp models.Basket
	if err := c.ShouldBindJSON(&agp); err != nil {
		log.Printf("Некорректный JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат данных"})
		return
	}
	err := services.PostBasket(agp)
	if err != nil {
		log.Printf("Ошибка при создании корзины: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при создании корзины"})
		return
	}
	log.Printf("Корзина c id_user %v и id_position %v успешно добавлена", agp.AuthUserID, agp.PositionID)
	c.JSON(http.StatusOK, agp)
}
func DeleteBasketHandler(c *gin.Context) {
	idParam := strings.TrimLeft(c.Param("id"), "/")
	if idParam == "" {
		log.Printf("ID не указан")
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID не указан"})
		return
	}

	var id int
	var isUserID bool
	var err error
	var message string

	if strings.HasPrefix(idParam, "user/") { //Удаление всей корзины
		idStr := strings.TrimPrefix(idParam, "user/")
		id, err = strconv.Atoi(idStr)
		isUserID = true
		message = "Корзина успешно очищена"
	} else { //Удаление позиции в корзине
		id, err = strconv.Atoi(idParam)
		isUserID = false
		message = "Позиция в корзине успешно удалена"
	}

	if err != nil {
		log.Printf("Неверный формат ID: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при конвертации id"})
		return
	}

	err = services.DeleteBasket(c, id, isUserID)

	if err != nil {
		log.Printf("Ошибка при удалении с id %v: %v", id, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Удаление с id %d выполнено успешно", id)
	c.JSON(http.StatusOK, gin.H{"message": message})
}

//func PutBasketHandler(c *gin.Context) {
//
//}
