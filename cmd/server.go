package main

import (
	"github.com/yeo/betterdev.link/baja"
	"log"
	"os"
)

func serve() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Cannot fetch current dir", err)
		return
	}

	addr := "0.0.0.0:1605"
	log.Println("Run server on", addr)
	baja.Serve(cwd, addr)
}
