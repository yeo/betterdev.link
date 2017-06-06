package main

import (
	"github.com/yeospace/better-dev.link/baja"
	"log"
	"os"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Cannot fetch current dir", err)
		return
	}

	addr := "127.0.0.1:1605"
	log.Println("Run server on", addr)
	baja.Serve(cwd, addr)
}
