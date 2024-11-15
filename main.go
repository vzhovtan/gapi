package main

import (
	"log"

	"gapi/internal/db"
	"gapi/internal/snip"
	"gapi/internal/transport"
)

func main() {
	d, err := db.New()
	if err != nil {
		log.Fatal(err)
	}
	svc := snip.NewService(d)
	server := transport.NewServer(svc)

	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}
}
