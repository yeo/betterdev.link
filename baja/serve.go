package baja

import (
	"github.com/yeospace/better-dev.link/baja/server"
)

func Serve(addr string) {
	server.Run(addr)
}
