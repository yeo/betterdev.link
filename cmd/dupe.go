package main

import (
	"fmt"
	"github.com/yeo/betterdev.link/baja"
)

func reportDupe(link string, previous, dupe int) {
	fmt.Println("Link", link, "appread in", previous, "and appear again in", dupe)
}

func detectDupe() {
	baja.CheckDupe(reportDupe)
}
