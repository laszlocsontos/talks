package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type timeHandler struct{}

func (th timeHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	currentTime := time.Now()

	fmt.Fprint(w, currentTime.Format(time.RFC1123Z))
}

func main() {
	err := http.ListenAndServe("localhost:4000", timeHandler{})

	if err != nil {
		log.Fatal(err)
	}
}
