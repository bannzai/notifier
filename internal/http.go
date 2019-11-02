package internal

import (
	"log"
	"net/http"

	"github.com/bannzai/notifier/pkg/driver"
)

func GitHub(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/github" {
		return
	}

	if r.Method != http.MethodPost {
		return
	}

	if err := driver.NewGitHub().Drive(r); err != nil {
		log.Printf("GitHub driver error with %v", err)
	}
}
