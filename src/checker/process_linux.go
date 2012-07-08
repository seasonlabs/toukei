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

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	fis, err := file.Readdir(0)
	if err != nil {
		log.Fatal(err)
	}

	ch = make(chan Stat)

	for _, fi := range fis {
		if fi.IsDir() {
			err = watcher.Watch(path + "/" + fi.Name() + "/refs/heads/master")
			if err != nil {
				log.Fatal(err)
			}
		}
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