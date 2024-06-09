package handler

import (
	"log/slog"
	"net/http"

	"github.com/dexguitar/chatapp/internal/queue"
	"github.com/gorilla/websocket"
)

type WSHandler struct {
	Hub queue.Hub
	*websocket.Upgrader
}

func NewWSHandler(hub queue.Hub) *WSHandler {
	return &WSHandler{
		Upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		Hub: hub,
	}
}

func (wsh *WSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("username")

	conn, err := wsh.Upgrade(w, r, nil)
	if err != nil {
		slog.Error("error %s when upgrading connection to websocket", err)
		http.Error(w, "error upgrading connection to websocket", http.StatusInternalServerError)
		return
	}

	wsh.Hub.AddConn(username, conn)
}
