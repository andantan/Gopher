package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	BUFFER_SIZE      int    = 1024
	MESSAGE_TYPE     int    = 1
	WS_URL           string = "/ws"
	WS_ORDERBOOK_URL string = "/ws.order"
	WS_PORT          string = ":8080"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  BUFFER_SIZE,
	WriteBufferSize: BUFFER_SIZE,
}

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) Orderbook(ws *websocket.Conn) {
	fmt.Println("New incoming connection from client to orderbook feed:", ws.RemoteAddr())

	for {
		payload := fmt.Sprintf("orderbook data -> %d", time.Now().UnixNano())

		if err := ws.WriteMessage(MESSAGE_TYPE, []byte(payload)); err != nil {
			fmt.Println("Write error:", err)
		}

		time.Sleep(time.Second)
	}
}

func (s *Server) handleWS(w http.ResponseWriter, r *http.Request) {
	ws := convertToWS(w, r)

	s.conns[ws] = true

	s.readLoop(ws)
}

func (s *Server) handleWSOrderbook(w http.ResponseWriter, r *http.Request) {
	ws := convertToWS(w, r)

	s.conns[ws] = true

	s.Orderbook(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	for {
		_, p, err := ws.ReadMessage()

		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Println("Read Error:", err)
			continue
		}

		fmt.Printf("[%s]: message(%s)\n",
			ws.RemoteAddr().String(), string(p))

		prefix := fmt.Sprintf("[%s]: %s\n", ws.RemoteAddr().String(), string(p))

		s.boradcast([]byte(prefix))
	}
}

func (s *Server) boradcast(p []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if err := ws.WriteMessage(MESSAGE_TYPE, p); err != nil {
				fmt.Println("Write error:", err)
			}
		}(ws)
	}
}

func convertToWS(w http.ResponseWriter, r *http.Request) (ws *websocket.Conn) {
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
	}

	log.Println("New incoming connection from client:", ws.RemoteAddr())

	return
}

func main() {
	server := NewServer()

	http.HandleFunc(WS_URL, server.handleWS)
	http.HandleFunc(WS_ORDERBOOK_URL, server.handleWSOrderbook)
	http.ListenAndServe(WS_PORT, nil)
}
