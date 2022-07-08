/*
based on these:
https://docs.microsoft.com/en-us/azure/azure-functions/create-first-function-vs-code-other?tabs=go%2Clinux#create-and-build-your-function
https://github.com/michalswi/simple-web-server
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	message := "This HTTP request triggered function executed successfully. Pass a 'name' in the query string for a personalized response.\n"
	name := r.URL.Query().Get("name")
	if name != "" {
		message = fmt.Sprintf("Hello, %s. This HTTP request triggered function executed successfully.\n", name)
	}
	fmt.Fprint(w, message)
}

func main() {
	port := os.Getenv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	http.HandleFunc("/api/simplewebserver", helloHandler)
	log.Println("Server is ready to handle requests at port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
