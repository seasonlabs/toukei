package checker

import (
	"log"
        "os"
       	"encoding/json"
       	"time"
       	"../commands"
)

//import "github.com/simonz05/godis"

type Stat struct {
	Path string
	Lines int
	Commits int
}

var ch chan Stat

func Check(path string) {
	for {
		process(path)
		time.Sleep(60 * time.Second)
	}
}

func process(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	fis, err := file.Readdir(0)
	if err != nil {
		log.Fatal(err)
	}

	ch = make(chan Stat)
	var dirs []os.FileInfo

	for _, fi := range fis {
		if fi.IsDir() {
			dirs = append(dirs, fi)
		}
	}

	for i, dir := range dirs {
		go func(current string, i int) {
			lines, err := commands.CountLines(current)
			if err != nil {
				log.Fatal(err)
			}

			commits, err := commands.CountCommits(current) 
			if err != nil {
				log.Fatal(err)
			}

	  		stat := Stat{Path: current, Lines: lines, Commits: commits}
	  		ch <- stat
	  	}(path + string(os.PathSeparator) + dir.Name(), i)
	}

	var stats []Stat
	for i := 0; i < len(dirs); i++ {
		stat := <-ch
		stats = append(stats, stat)
	}

	statsJson, err := json.Marshal(stats)
	if err != nil {
	    log.Fatal(err)
	}
	
	return string(statsJson)
}