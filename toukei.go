package main

import (
	"fmt"
	"log"
        "os"
)

func main() {
	if err := os.Chdir("."); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Lines: %d\n", countLines())
  	fmt.Printf("Commits: %d\n", countCommits())
}