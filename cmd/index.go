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

	indexer := &baja.Indexer{
		AlgoliaAppID: os.Getenv("ALGOLIA_APP_ID"),
		AlgoliaKey:   os.Getenv("ALGOLIA_KEY"),
	}

	indexer.Index(cwd)
}
