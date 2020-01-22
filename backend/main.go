package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func status(w http.ResponseWriter, r *http.Request) {
	// Arbitrary sleep so that we can demonstrate autoscaler
	log.Println("requests received")
	time.Sleep(101 * time.Millisecond)
	fmt.Fprintln(w, "status")
}

func home(w http.ResponseWriter, r *http.Request) {
	// Arbitrary sleep so that we can demonstrate autoscaler
	log.Println("requests received")
	time.Sleep(101 * time.Millisecond)
	fmt.Fprintln(w, "Hi I'm the backend")
}

func main() {
	log.SetPrefix("")
	log.SetFlags(log.Lshortfile)
	log.Println("init started")

	http.HandleFunc("/", home)
	http.HandleFunc("/status", status)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

