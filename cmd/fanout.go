package main

import (
	"fmt"
	"github.com/yeo/betterdev.link/baja"
)

func fanout(issue, env string) {
	fmt.Printf("Fanout %s\n", issue)
	baja.Fanout(issue, env)
}
