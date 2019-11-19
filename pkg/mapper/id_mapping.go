package mapper

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/bannzai/notifier/pkg/sender"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type (
	GitHub struct {
		Login string
	}
	Slack struct {
		ID   string
		Name string
	}
)

type User struct {
	ID string
	GitHub
	Slack
}

func fetchUsers() ([]User, error) {
	path := os.Getenv("YAML_FILE_PATH")
	if url, err := url.ParseRequestURI(path); err != nil {
		response, err := http.Get(url.String())
		if err != nil {
			return []User{}, errors.Wrapf(err, "http error with url: %s", url.String())
		}
		defer response.Body.Close()

		body := response.Body
		bytes, err := ioutil.ReadAll(body)
		if err != nil {
			return []User{}, errors.Wrapf(err, "ioutil.ReadAll error with http response body %v", body)
		}
		mapping := &[]User{}
		if err := yaml.Unmarshal(bytes, mapping); err != nil {
			return []User{}, errors.Wrapf(err, "Decode error to yaml from %s", string(bytes))
		}
		return *mapping, nil
	} else {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return []User{}, errors.Wrapf(err, "ioutil.ReadFile error with path: %s", path)
		}
		mapping := &[]User{}
		if err := yaml.Unmarshal(data, mapping); err != nil {
			return []User{}, errors.Wrapf(err, "yaml.Unmarshal error with %s", string(data))
		}
		return *mapping, nil
	}
}

func extractUserFromGitHub(
	users []User,
	githubUserName string,
	extractedContentType sender.ContentType,
) (Slack, bool) {
	switch extractedContentType {
	case sender.SlackContentType:
		for _, user := range users {
			if user.GitHub.Login == githubUserName {
				return user.Slack, true
			}
		}
		return Slack{}, false
	default:
		return Slack{}, false
	}
}
