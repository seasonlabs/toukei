package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"toukei/checker"
)

import "github.com/kylelemons/go-gypsy/yaml"
import "github.com/simonz05/godis"
import "code.google.com/p/go.net/websocket"

var port *int = flag.Int("p", 8080, "Port to listen.")

//func hello(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Hello")
//}

func main() {
	flag.Parse()
	config, err := yaml.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("readfile(%q): %s", "config.yml", err)
	}

	path, err := config.Get("path")
	if err != nil {
		log.Fatal(err)
	}

	go checker.Check(path)
	go listen()

	http.Handle("/json", websocket.Handler(jsonServer))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", MainServer)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func listen() {
	c := godis.New("", 0, "")

	s, err := c.Subscribe("toukei")
	if err != nil {
		log.Fatal(err)
	}

	for {
		m := <-s.Messages
		println(m.Elem.String())
	}
}

func MainServer(w http.ResponseWriter, req *http.Request) {
	t := template.Must(template.New("foo").ParseGlob("index.html"))
	if err := t.ExecuteTemplate(w, "index", req.Host+":"+req.URL.Scheme); err != nil {
		log.Fatal(err)
	}
}

// jsonServer handles listening to reddis messages and push the result to connected clients 
func jsonServer(ws *websocket.Conn) {
	c := godis.New("", 0, "")

	elem, err := c.Get("toukei")
	if err != nil {
		log.Println(err)
		elem = []byte("{}")
	}
	println(string(elem))
	websocketSend(ws, elem)

	s, err := c.Subscribe("toukei")
	if err != nil {
		log.Fatal(err)
	}

	for {
		m := <-s.Messages
		websocketSend(ws, m.Elem)
	}
}

func websocketSend(ws *websocket.Conn, elem []byte) {
	var stat checker.Stat

	if err := json.Unmarshal(elem, &stat); err != nil {
		log.Fatal(err)
	}

	// Send send a text message serialized T as JSON.
	if err := websocket.JSON.Send(ws, stat); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("send: %#v\n", stat)
}