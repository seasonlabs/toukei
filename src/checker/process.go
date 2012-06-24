package checker

import (
	"log"
        "os"
       	"encoding/json"
       	"time"
       	"../commands"
)

import "github.com/simonz05/godis"

type Stat struct {
	Path string
	Lines int
	Commits int
	Repos int
}

var ch chan Stat

func Check(path string) {
	for {
		publish(process(path))
		time.Sleep(20 * time.Second)
	}
}

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
			lines, _ := commands.CountLines(current)
			commits, _ := commands.CountCommits(current) 

	  		stat := Stat{Path: current, Lines: lines, Commits: commits}
	  		ch <- stat
	  	}(path + string(os.PathSeparator) + dir.Name(), i)
	}

	stats := Stat{Path: "Total", Lines: 0, Commits: 0, Repos: len(dirs)}
	for i := 0; i < len(dirs); i++ {
		stat := <-ch
		stats.Lines += stat.Lines
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

	if _, err := c.Publish("toukei", statsJson); err != nil {
		log.Fatal(err)
	}
}