package main

import (
	"github.com/yeospace/better-dev.link/baja"
	"log"
	//"os"
)

func main() {
	addr := "127.0.0.1:1605"
	log.Println("Run server on", addr)
	baja.Serve(addr)
}
