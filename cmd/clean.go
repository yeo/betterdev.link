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

	baja.Clean(cwd)
}
