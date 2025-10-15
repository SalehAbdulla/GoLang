package main

import (
	"fmt"
	"log"
	"net/http"
)

type Handlers struct{}

func (h *Handlers) SendHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from docker file</h1>")
}

func main() {
	h := Handlers{}
	http.HandleFunc("/", h.SendHello)
	log.Fatal(http.ListenAndServe(":5173", nil))
}
