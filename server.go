package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/receive", func(w http.ResponseWriter, r *http.Request) {
		io.Copy(ioutil.Discard, r.Body)
	})
	log.Fatal(http.ListenAndServe(":1234", nil))
}