package internal

import (
	"log"
	"net/http"
	"os"

	"github.com/bannzai/notifier/pkg/driver"
	"github.com/bannzai/notifier/pkg/mapper"
	"github.com/bannzai/notifier/pkg/parser"
	"github.com/bannzai/notifier/pkg/sender"
	"github.com/pkg/errors"
)

func GitHub(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/github" {
		return
	}

	if r.Method != http.MethodPost {
		return
	}

	// TODO: Flexing with command line argument
	githubToSlackDriver := driver.New(
		parser.NewGitHub(),
		sender.NewSlack(
			os.Getenv("NOTIFIER_SLACK_TOKEN"),
			mapper.New(),
		),
	)
	if err := githubToSlackDriver.Drive(r); err != nil {
		log.Printf("GitHub driver error with %v", err)
		if err = errors.Cause(err); err != nil {
			// TODO: Refactor for integrate error messages with original error type
			log.Printf("GitHub driver error from %v", err)
		}
	}

	log.Println("successfully post message to slack")
}
