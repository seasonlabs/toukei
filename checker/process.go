package checker

import (
	"log"
        "os"
       	"encoding/json"
       	"github.com/seasonlabs/toukei/commands"
)

import "github.com/simonz05/godis"

type Stat struct {
	Path string
	Files int
	Commits int
	Repos int
}

var ch chan Stat

func process(path string) Stat {
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
			files, _ := commands.CountFiles(current)
			commits, _ := commands.CountCommits(current) 

	  		stat := Stat{Path: current, Files: files, Commits: commits}
	  		ch <- stat
	  	}(path + string(os.PathSeparator) + dir.Name(), i)
	}

	stats := Stat{Path: "Total", Files: 0, Commits: 0, Repos: len(dirs)}
	for i := 0; i < len(dirs); i++ {
		stat := <-ch
		stats.Files += stat.Files
		stats.Commits += stat.Commits
	}

	return stats
}

func publish(stats Stat) {
	c := godis.New("", 0, "")

	statsJson, err := json.Marshal(stats)
	if err != nil {
	    log.Fatal(err)
	}

	if err := c.Set("toukei", statsJson); err != nil {
		log.Fatal(err)
	}

	if _, err := c.Publish("toukei", statsJson); err != nil {
		log.Fatal(err)
	}
}