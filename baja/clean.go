package baja

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func Clean(action, dir string) {
	log.Println("Clean", dir, "public")
	// os.RemoveAll(dir + "/public")

	contacts, err := ioutil.ReadFile(fmt.Sprintf("./content/%s.csv", "prod"))
	if err != nil {
		log.Fatal("Cannot read file", err)
	}

	r := csv.NewReader(bytes.NewReader(contacts))

	file, err := os.Open(fmt.Sprintf("./content/%s.csv", action))
	if err != nil {
		log.Fatal("Error when parse click file")
	}
	defer file.Close()

	var clickedUsers map[string]bool
	clickedUsers = make(map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		clickedUsers[scanner.Text()] = true
	}

	count := 0
	for {
		count = count + 1
		record, err := r.Read()
		if count <= 1 {
			continue
		}
		if err != nil {
			break
		}

		h := sha256.New()
		h.Write([]byte(record[0]))
		emailHash := fmt.Sprintf("%x", h.Sum(nil))

		if _, ok := clickedUsers[emailHash]; ok == false {
			fmt.Println(record[0])
		}
	}
}
