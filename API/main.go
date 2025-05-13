package main

import (
	"flag"
	"fmt"
	"log"
	"opet/API/api"
	"opet/API/storage"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "the server address")
	flag.Parse()

	storage := storage.NewMemoryStorage()

	server := api.NewServer(*listenAddr, storage)
	fmt.Println("server running on port:", *listenAddr)
	log.Fatal(server.Start())
}
