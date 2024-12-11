package handlers

import (
	"chat-app/db"
	"chat-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateChat(c *gin.Context) {
	var chat models.Chat
	if err := c.ShouldBindJSON(&chat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&chat).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create chat"})
		return
	}

	c.JSON(http.StatusCreated, chat)
}

func GetMessages(c *gin.Context) {
	chatID := c.Param("chatID") // Получаем ID чата из маршрута

	var messages []models.Message                      // Срез для хранения сообщений
	db.DB.Where("chat_id = ?", chatID).Find(&messages) // Запрос к базе данных

	c.JSON(http.StatusOK, messages) // Возвращаем сообщения в формате JSON
}

func SendMessage(c *gin.Context) {
	var message models.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&message).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send message"})
		return
	}

	c.JSON(http.StatusCreated, message)
}
