package main

import (
	"github.com/yeospace/better-dev.link/baja"
	"log"
	"os"
)

func main() {
	dest := os.Getenv("BAJA_DEPLOY_TO")
	host := os.Getenv("BAJA_DEPLOY_HOST")
	log.Println("Deploy to", host, dest)
	baja.DeploySSH(host, dest)
}
