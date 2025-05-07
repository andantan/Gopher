package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

type Message struct {
	from    string
	payload []byte
}

type Server struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
	msgch      chan *Message
	peerMap    map[net.Conn]bool
}

func NewServer(listenaddr string) *Server {
	return &Server{
		listenAddr: listenaddr,
		quitch:     make(chan struct{}),
		msgch:      make(chan *Message, 10),
		peerMap:    make(map[net.Conn]bool),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)

	if err != nil {
		return err
	}

	defer ln.Close()

	s.ln = ln

	go s.acceptLoop()

	<-s.quitch

	close(s.msgch)

	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()

		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}

		s.registerConn(&conn)

		fmt.Println("New connection:", conn.RemoteAddr().String())

		s.msgch <- &Message{
			from:    conn.RemoteAddr().String(),
			payload: fmt.Appendf(nil, "***New connection***\n\r"),
		}

		go s.readLoop(conn)
	}
}

func (s *Server) registerConn(conn *net.Conn) {
	s.peerMap[*conn] = true

	s.notifyConnected(conn)
}

func (s *Server) notifyConnected(conn *net.Conn) {
	msg := fmt.Sprintf("Your ID: %s\n\r", (*conn).RemoteAddr().String())

	if _, err := (*conn).Write([]byte(msg)); err != nil {
		log.Fatal(err)
	}
}

func (s *Server) removeConn(conn *net.Conn) {
	delete(s.peerMap, *conn)
}

func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()

	r := make([]byte, 1)
	buf := make([]byte, 2048)

	for {
		if _, err := conn.Read(r); err != nil {
			if err == io.EOF {
				s.removeConn(&conn)

				fmt.Println("Close connection:", conn.RemoteAddr().String())

				s.msgch <- &Message{
					from:    conn.RemoteAddr().String(),
					payload: fmt.Appendf(nil, "***User disconnected***\n\r"),
				}

				break
			}

			fmt.Println("Read error:", err)
			continue
		}

		buf = append(buf, r...)

		if r[0] == 10 {
			s.msgch <- &Message{
				from:    conn.RemoteAddr().String(),
				payload: buf,
			}

			fmt.Printf("[%s]: %s", conn.RemoteAddr().String(), string(buf))

			buf = buf[:0]
		}
	}

}

func (s *Server) broadcast() {
	for msg := range s.msgch { // Stand-by
		for conn, connected := range s.peerMap { // Broadcasting
			if connected && msg.from != conn.RemoteAddr().String() {
				data := fmt.Sprintf("[%s]: %s", msg.from, string(msg.payload))

				if _, err := conn.Write([]byte(data)); err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

func main() {
	server := NewServer(":3000")

	go server.broadcast()

	log.Fatal(server.Start())
}
