package baja

import (
	"github.com/yeo/betterdev.link/baja/server"
)

func Serve(cwd, addr string) {
	watcher := NewWatcher(cwd)
	go watcher.Run()
	server.Run(addr)
}
