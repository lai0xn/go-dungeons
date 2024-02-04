package main

import (
	"flag"

	"github.com/jn0x/reddigo/http/server"
	"github.com/jn0x/reddigo/storage"
)

func main() {
	port := flag.String("port", ":5000", "-p")
	flag.Parse()
	server := server.NewServer(*port)
	storage.Connect()
	server.Serve()
}
