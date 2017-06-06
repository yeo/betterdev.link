package baja

import (
	"github.com/yeospace/better-dev.link/baja/server"
)

func Serve(cwd, addr string) {
	watcher := NewWatcher(cwd)
	go watcher.Run()
	server.Run(addr)
}
