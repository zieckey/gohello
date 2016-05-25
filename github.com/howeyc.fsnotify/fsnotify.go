package main

import (
	"log"
	"os"
	"github.com/howeyc/fsnotify"
)

// IsDir returns true if given path is a directory,
// or returns false when it's a file or does not exist.
func IsDir(dir string) bool {
	f, e := os.Stat(dir)
	if e != nil {
		return false
	}
	return f.IsDir()
}

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)

	// Process events
	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				log.Println("event:", ev, " name=", ev.Name)
				if ev.IsCreate() && IsDir(ev.Name) {
					watcher.Watch(ev.Name)
					//TODO if we renamed ev.Name laterly, we should add the new name to the watching list
					
				}
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Watch("e:/1")
	//err = watcher.Watch("/tmp/a.txt")
	if err != nil {
		log.Fatal(err)
	}

	<-done

	/* ... do stuff ... */
	watcher.Close()
	
}