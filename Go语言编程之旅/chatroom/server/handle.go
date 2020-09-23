package server

import (
	"net/http"
)

func RegisterHandle() {
	// 广播消息
	//go logic.Broadcaster.Start()

	http.HandleFunc("/", homeHandleFunc)
	//http.HandleFunc("/ws", WebSocketHanldeFunc)
}
