package checker

import (
	"time"
	"../inotify"
	"log"
)

func Check(path string) {
	watcher, err := inotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	
	err = watcher.Watch(path)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case ev := <-watcher.Event:
	    		log.Println("event:", ev)
		case err := <-watcher.Error:
	    		log.Println("error:", err)
		}
	}
}