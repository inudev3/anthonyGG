package main

import (
	"bytes"
	"fmt"
	"io"
)

type Conn struct {
	io.Writer
}
type Server struct {
	peers map[*Conn]bool
}

func (c *Conn) Write(p []byte) (int, error) {
	fmt.Println("writing to underlying connection:", string(p))
	return c.Writer.Write(p)
}
func NewConn() *Conn {
	return &Conn{
		Writer: new(bytes.Buffer),
	}
}
func NewServer() *Server {
	s := &Server{peers: make(map[*Conn]bool)}
	for i := 0; i < 10; i++ {
		s.peers[NewConn()] = true
	}
	return s
}
func (s *Server) broadcast(msg []byte) error {
	peers := []io.Writer{}

	for peer := range s.peers {
		peers = append(peers, peer)
	}
	mw := io.MultiWriter(peers...)
	_, err := mw.Write(msg)
	return err
}
func main() {
	s := NewServer()

	s.broadcast([]byte("foo"))
}
