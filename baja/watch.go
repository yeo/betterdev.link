package baja

import (
	"fmt"
	"github.com/radovskyb/watcher"
	"log"
	"time"
)

type Watcher struct {
	cwd     string
	watcher *watcher.Watcher
}

func NewWatcher(cwd string) *Watcher {
	w := watcher.New()
	return &Watcher{
		cwd:     cwd,
		watcher: w,
	}
}

func (w *Watcher) Run() {
	log.Println("Precompile")
	Compile(w.cwd)

	directories := [2]string{"./content", "./themes"}
	for _, d := range directories {
		if err := w.watcher.AddRecursive(d); err != nil {
			log.Fatalln(err)
		}
	}

	go func() {
		for {
			select {
			case event := <-w.watcher.Event:
				fmt.Println("Receiv:", event) // Print the event's info.
				Compile(w.cwd)
			case err := <-w.watcher.Error:
				log.Fatalln(err)
			case <-w.watcher.Closed:
				return
			}
		}
	}()

	if err := w.watcher.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}
