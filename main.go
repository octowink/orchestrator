package main

import (
	"flag"
	"log"
	"net/http"
  "github.com/gorilla/websocket"
)

var addr = flag.String("addr", ":9090", "http service address")

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "index.html")
}

type serverMessage struct {
  Type int `json:"type"`
  Message string `json:"message,omitempty"`
}

func main() {
  flag.Parse()
  http.HandleFunc("/", serveHome)

  upgrader.CheckOrigin = func(r *http.Request) bool { return true }

  http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
    _, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
      log.Println(err)
      return
    }

    log.Println("TODO add listeners")
  })

  err := http.ListenAndServe(*addr, nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
