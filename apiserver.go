package main

import (
	"log"

	"github.com/Vasilesk/go-backend-cybersport/apiserver/server"
)

func main() {
	log.Fatal(server.RunHTTPServer(":3003"))
}
