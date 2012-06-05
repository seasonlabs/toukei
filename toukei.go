package main

import (
	"fmt"
	"log"
        "os"
)

func main() {
	if err := os.Chdir("../savia"); err != nil {
		log.Fatal(err)
	}

  	fmt.Printf("%d\n", countLines())
}