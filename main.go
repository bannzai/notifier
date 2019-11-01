package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const defaultPort = "5000"

func main() {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = defaultPort
	}

	fmt.Printf("Start process!! port :%s", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
