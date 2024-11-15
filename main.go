package main

import (
	"log"

	"github.com/vzhovtan/gapi/internal/db"
	"github.com/vzhovtan/gapi/internal/snip"
	"github.com/vzhovtan/gapi/internal/transport"
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
