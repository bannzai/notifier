package internal

import (
	"log"
	"net/http"

	"github.com/bannzai/notifier/pkg/drivers/github"
)

func GitHub(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/github" {
		return
	}

	if r.Method != http.MethodPost {
		return
	}

	if err := github.NewDriver().Drive(r); err != nil {
		log.Printf("GitHub driver error with %v", err)
	}
}
