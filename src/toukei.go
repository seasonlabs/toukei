package main

import (
	"fmt"
	"log"
       	"net/http"
       	"./checker"
)

import "github.com/kylelemons/go-gypsy/yaml"
//import "github.com/simonz05/godis"

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func main() {
	config, err := yaml.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("readfile(%q): %s", "config.yml", err)
	}

	path, err := config.Get("path")
	if err != nil {
		log.Fatal(err)
	}

	go checker.Check(path)

	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":8080", nil))

}