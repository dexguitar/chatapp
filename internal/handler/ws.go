package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/dexguitar/chatapp/internal/queue"
	"github.com/gorilla/websocket"
)

type WSHandler struct {
	Hub            queue.Hub
	MessageService MessageService
	*websocket.Upgrader
}

func NewWSHandler(hub queue.Hub, ms MessageService) *WSHandler {
	return &WSHandler{
		Upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		Hub:            hub,
		MessageService: ms,
	}
}

func (wsh *WSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := wsh.Upgrade(w, r, nil)
	if err != nil {
		slog.Error(fmt.Sprintf("WebSocket upgrade failed: %v", err))
		return
	}
	defer conn.Close()

	userID := r.URL.Query().Get("userID")
	if userID == "" {
		slog.Error("userID not provided")
		conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "userID required"))
		return
	}

	wsh.Hub.AddConn(userID, conn)
	defer wsh.Hub.RemoveConn(userID)

	slog.Info(fmt.Sprintf("User `%s` connected", userID))

	for {
		if _, _, err := conn.NextReader(); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				slog.Error(fmt.Sprintf("Unexpected close error: %v", err))
			}
			break
		}
	}
}
