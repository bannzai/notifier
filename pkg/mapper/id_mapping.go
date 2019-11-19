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

type IDMapping struct {
	GitHubList []GitHub
	SlackList  []Slack
}

func fetchIDMap() (IDMapping, error) {
	path := os.Getenv("YAML_FILE_PATH")
	if url, err := url.ParseRequestURI(path); err != nil {
		response, err := http.Get(url.String())
		if err != nil {
			return IDMapping{}, errors.Wrapf(err, "http error with url: %s", url.String())
		}
		defer response.Body.Close()

		body := response.Body
		bytes, err := ioutil.ReadAll(body)
		if err != nil {
			return IDMapping{}, errors.Wrapf(err, "ioutil.ReadAll error with http response body %v", body)
		}
		mapping := &IDMapping{}
		if err := yaml.Unmarshal(bytes, mapping); err != nil {
			return IDMapping{}, errors.Wrapf(err, "Decode error to yaml from %s", string(bytes))
		}
		return *mapping, nil
	} else {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return IDMapping{}, errors.Wrapf(err, "ioutil.ReadFile error with path: %s", path)
		}
		mapping := &IDMapping{}
		if err := yaml.Unmarshal(data, mapping); err != nil {
			return IDMapping{}, errors.Wrapf(err, "yaml.Unmarshal error with %s", string(data))
		}
		return *mapping, nil
	}
}

func (mapper IDMapping) extractFromGitHub(username string, extractedContentType sender.ContentType) string {
	switch extractedContentType {
	case sender.SlackContentType:
		var matchedContent GitHub
		for _, github := range mapper.GitHubList {
			if username == github.Login {
				matchedContent = github
			}
		}
		for _, slack := range mapper.SlackList {
			if username == slack.ID {
				return username
			}
			if username == slack.Name {
				return username
			}
		}
		return ""
	default:
		return ""
	}
}
