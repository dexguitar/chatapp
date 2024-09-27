package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/dexguitar/chatapp/internal/model"
	"github.com/dexguitar/chatapp/internal/queue"
	"github.com/gorilla/websocket"
)

type WSHandler struct {
	Hub   queue.Hub
	Queue *queue.Queue
	*websocket.Upgrader
}

func NewWSHandler(hub queue.Hub, q *queue.Queue) *WSHandler {
	return &WSHandler{
		Upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		Hub:   hub,
		Queue: q,
	}
}

func (wsh *WSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := wsh.Upgrade(w, r, nil)
	if err != nil {
		slog.Error(fmt.Sprintf("WebSocket upgrade failed: %v", err))
		return
	}
	defer conn.Close()

	username := r.URL.Query().Get("username")
	if username == "" {
		slog.Error("Username not provided")
		conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Username required"))
		return
	}

	wsh.Hub.AddConn(username, conn)
	defer wsh.Hub.RemoveConn(username)

	slog.Info(fmt.Sprintf("User `%s` connected", username))

	for {
		var msg = model.Message{Username: username}
		err := conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				slog.Error(fmt.Sprintf("Unexpected close error: %v", err))
			}
			break
		}

		// msg.Timestamp = time.Now()
		// msg.ID = generateUniqueID()

		err = wsh.Queue.WriteMessage(r.Context(), &msg)
		if err != nil {
			slog.Error(fmt.Sprintf("Failed to write message to Kafka: %v", err))
			conn.WriteJSON(map[string]string{
				"error": "Failed to send message",
			})
			return
		}
	}
}
