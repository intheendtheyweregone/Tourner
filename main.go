package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/intheendtheyweregone/tourner/sharex"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {

	// Denies request if not a POST Request.
	if r.Method != "POST" {
		fmt.Fprintf(w, "This endpoint requires a POST Request. \n")
		return
	}

	// Reading File
	file, handler, err := r.FormFile("files[]")
	if err != nil {
		fmt.Fprintf(w, "There was an error: %v.", err)

		return
	}
	defer file.Close()

	// Calls method to randomly select site
	siteType, site := selectSite()

	body, err := sharex.ProcessUpload(file, handler, site, siteType)
	if err != nil {
		fmt.Fprintf(w, "There was an error: %v", err)
		return
	}

	response := string(body)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%v", response)
}

func main() {
	// Handle CLI Flags
	portArgument := handleCLI()

	// Create route for uploading
	http.HandleFunc("/upload", uploadHandler)

	fmt.Printf("[Tourner] Started successfully on port %v", *portArgument)
	// Start web server
	log.Fatal(http.ListenAndServe("127.0.0.1:"+*portArgument, nil))
}
