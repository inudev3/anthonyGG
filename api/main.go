package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "the server address")
	storage := &MongoDb{}
	server := Newserver(*listenAddr, storage)
	fmt.Println("server running on port:", *listenAddr)
	log.Fatal(server.Start())
	flag.Parse()
}
