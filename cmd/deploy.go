package main

import (
	"github.com/yeo/betterdev.link/baja"
	"log"
	"os"
)

func deploy() {
	dest := os.Getenv("BAJA_DEPLOY_TO")
	host := os.Getenv("BAJA_DEPLOY_HOST")
	log.Println("Deploy to", host, dest)
	baja.DeploySSH(host, dest)
}
