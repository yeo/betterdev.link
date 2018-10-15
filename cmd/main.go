package main

import (
	"fmt"
	"github.com/yeo/betterdev.link/baja"
	"log"
	"os"
)

var (
	Version   string
	GitCommit string
)

func main() {
	fmt.Printf("BetterDev %s Build %s\n", Version, GitCommit)

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Cannot fetch current dir", err)
		return
	}

	if len(os.Args) == 1 {
		log.Println("-> Compile")
		baja.Compile(cwd)
		return
	}

	switch os.Args[1] {
	case "build":
		baja.Compile(cwd)
	case "clean":
		clean(os.Args[2])
	case "serve", "server":
		serve()
	case "dupe":
		detectDupe()
	case "fanout":
		fanout(os.Args[2], os.Args[3])
	}
}
