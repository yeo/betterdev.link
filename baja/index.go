package baja

import (
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"log"
)

type Indexer struct {
	AlgoliaAppID string
	AlgoliaKey   string
}

func (i *Indexer) Index(dir string) {
	client := algoliasearch.NewClient(i.AlgoliaAppID, i.AlgoliaKey)
	index := client.InitIndex("links")

	page := buildPage()

	var objects []algoliasearch.Object

	for _, issue := range page.Issues {
		for _, link := range issue.Links {
			objects = append(objects, algoliasearch.Object{
				"ObjectID":    issue.Name + link.URI,
				"Title":       link.Title,
				"Description": link.Description,
			})
			log.Println("ObjectID", issue.Name, link.URI)
		}
	}

	res, err := index.AddObjects(objects)
	if err != nil {
		log.Fatal("Index Error", err)
	}
	log.Println("indes success", res)
}
