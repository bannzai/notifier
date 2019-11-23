package internal

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/bannzai/notifier/pkg/driver"
	"github.com/bannzai/notifier/pkg/logger"
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

	response := Response{}
	if err := githubToSlackDriver.Drive(r); err != nil {
		logger.Logf("GitHub driver error with %v", err)
		response = Response{
			Success: false,
			Errors:  []Errors{{Message: err.Error()}},
		}
	} else {
		logger.Log("successfully post message to slack")
		response = Response{
			Success: true,
			Errors:  []Errors{},
		}
	}

	bytes, err := write(response)
	if err != nil {
		logger.Logf("write respone error with %v", err)
	}
	w.Write(bytes)
}

func write(response Response) ([]byte, error) {
	bytes, err := json.Marshal(response)
	if err != nil {
		return []byte{}, errors.Wrap(err, "json.Marshal(response) is error")
	}
	return bytes, nil
}

type (
	Response struct {
		Success bool     `json:"success"`
		Errors  []Errors `json:"errors"`
	}
	Errors struct {
		Message string `json:"message"`
	}
)
