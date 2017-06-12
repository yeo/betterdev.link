package baja

import (
	"log"
	"os"
)

func Clean(dir string) {
	log.Println("Clean", dir, "public")
	os.RemoveAll(dir + "/public")
}
