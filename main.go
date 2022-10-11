package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/intheendtheyweregone/tourner/sharex"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		fmt.Fprintf(w, "This endpoint requires a POST Request. \n")
		return
	}

	file, handler, err := r.FormFile("files[]")
	if err != nil {
		fmt.Fprintf(w, "There was an error: %v.", err)

		return
	}
	defer file.Close()

	site := selectSite()
	body, err := sharex.HandleUpload(file, handler, site)
	if err != nil {
		fmt.Fprintf(w, "There was an error: %v", err)
		return
	}

	response := string(body)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%v", response)
}

func main() {

	http.HandleFunc("/upload", uploadHandler)

	fmt.Println("Started...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
