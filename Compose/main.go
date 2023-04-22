package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

type TransformFunc func(string) string
type Server struct {
	filenametransform TransformFunc
}

func (s *Server) handleRequest(filename string) error {
	newFilename := s.filenametransform(filename)
	fmt.Println("new filename:", newFilename)
	return nil
}
func hashFilename(filename string) string {
	hash := sha256.Sum256([]byte(filename))
	newFilename := hex.EncodeToString(hash[:])
	return newFilename
}

func prefixFilename(prefix string) TransformFunc {
	return func(s string) string {
		return prefix + s
	}
}
func main() {
	if len(os.Args) < 2 {
		log.Fatal("arg needed!")
	}
	prefix := os.Args[1]
	s := &Server{
		filenametransform: prefixFilename(prefix),
	}
	s.handleRequest("cool_picture.jpg")
}
