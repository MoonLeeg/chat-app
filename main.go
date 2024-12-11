package main

import (
	"chat-app/db"
	"chat-app/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.SetupDatabase()
	r := gin.Default()

	// Регистрация и аутентификация
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// Создание чата и отправка сообщения
	r.POST("/chats", handlers.CreateChat)
	r.POST("/messages", handlers.SendMessage)

	// Получение сообщений для конкретного чата
	r.GET("/chats/:chatID/messages", handlers.GetMessages)

	// WebSocket обработчик
	r.GET("/ws", handlers.GetMessages)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Chat App!",
		})
	})

	r.Run() // Запуск сервера на 8080 порту
}
