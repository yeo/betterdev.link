package baja

import (
	"log"
)

type Indexer struct {
	AlgoliaAppID string
	AlgoliaKey   string
}

func (i *Indexer) Index(dir string) {
	log.Println("TODO")
}
