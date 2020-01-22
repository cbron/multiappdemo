package main

import (
	"fmt"
	"log"
	"net/http"
    "os"
    "io/ioutil"
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
	fmt.Fprintln(w, "home\nbackend response:", backendResponse())

}

func backendResponse() (string) {
	endpoint := os.Getenv("BACKEND")
	if endpoint == "" {
		return "no backend specified"
	}
	resp, err := http.Get(endpoint)
	if err != nil {
		return "error"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "error"
	}
	return string(body)
}

func main() {
	log.SetPrefix("")
	log.SetFlags(log.Lshortfile)
	log.Println("init started")

	http.HandleFunc("/", home)
	http.HandleFunc("/status", status)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
