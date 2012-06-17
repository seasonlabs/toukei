package main

import (
	"fmt"
	"log"
        "os"
       	"net/http"
       	"encoding/json"
)

import "github.com/kylelemons/go-gypsy/yaml"
import "github.com/simonz05/godis/redis"

type Stat struct {
	Path string
	Lines int
	Commits int
}

var path string
var ch chan Stat

func process(w http.ResponseWriter, r *http.Request) {
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
			lines, err := countLines(current)
			if err != nil {
				log.Fatal(err)
			}

			commits, err := countCommits(current) 
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

	b, err := json.Marshal(stats)
	if err != nil {
	    log.Fatal(err)
	}
	
	fmt.Fprintf(w, string(b))
}

func main() {
	config, err := yaml.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("readfile(%q): %s", "config.yml", err)
	}

	path, err = config.Get("path")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", process)
	log.Fatal(http.ListenAndServe(":8080", nil))
}