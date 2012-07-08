package checker

import (
	"../inotify"
	"log"
	"os"
)

func Check(path string) {
	publish(process(path))

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
			err = watcher.Watch(path + "/" + fi.Name() + "/refs/heads")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	
	for {
		select {
		case ev := <-watcher.Event:
	    		log.Println("event:", ev)
	    		publish(process(path))
		case err := <-watcher.Error:
	    		log.Println("error:", err)
		}
	}
}