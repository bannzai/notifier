package parser

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type IDMapping struct {
	GitHub struct {
		Login string
	}
	Slack struct {
		ID   string
		Name string
	}
}

func FetchIDMap() (IDMapping, error) {
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
