package main

import (
	"flag"
	"log"

	"github.com/yeo/betterdev.link/baja"
)

func main() {
	var issue = flag.Int("issue", 0, "Issue #")
	flag.Parse()

	if *issue == 0 {
		log.Fatal("Invalid issue: ", *issue)
	}

	baja.Schedule(*issue)
}
