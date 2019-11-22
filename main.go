package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bannzai/notifier/internal"
	"github.com/bannzai/notifier/pkg/logger"
)

const defaultPort = "5000"

func main() {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = defaultPort
	}

	logger.Logf("Start process!! port: %s", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	http.HandleFunc("/github", internal.GitHub)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
