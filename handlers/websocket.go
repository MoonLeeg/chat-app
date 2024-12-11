package handlers

import (
	"chat-app/db"
	"chat-app/models"

	"github.com/dgrijalva/jwt-go"

	"golang.org/x/net/websocket"
)

var clients = make(map[*websocket.Conn]bool) // Все соединения клиентов

// Обработчик WebSocket соединений
func WebSocketHandler(ws *websocket.Conn) {
	// Извлечение запроса
	tokenString := ws.Request().URL.Query().Get("token")

	claims := &Claims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		ws.Close()
		return
	}

	clients[ws] = true
	defer func() {
		delete(clients, ws)
		ws.Close()
	}()

	for {
		var msg models.Message
		if err := websocket.JSON.Receive(ws, &msg); err != nil {
			break // Ошибка получения данных, выходим
		}

		// Сохраняем сообщение в базе данных
		if err := db.DB.Create(&msg).Error; err == nil {
			// Отправляем сообщение всем подключенным клиентам
			for client := range clients {
				if err := websocket.JSON.Send(client, msg); err != nil {
					client.Close()
					delete(clients, client)
				}
			}
		}
	}
}
