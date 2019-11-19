package mapper

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type (
	GitHub struct {
		Login string `yaml:"login"`
	}
	Slack struct {
		ID   string `yaml:"id"`
		Name string `yaml:"name"`
	}
)

type User struct {
	ID     string `yaml:"id"`
	GitHub `yaml:"github"`
	Slack  `yaml:"slack"`
}

func isRemoteURL(path string) bool {
	return strings.Contains(path, "http://") || strings.Contains(path, "https://")
}

func fetchUsers() ([]User, error) {
	path := os.Getenv("YAML_FILE_PATH")
	if isRemoteURL(path) {
		response, err := http.Get(path)
		if err != nil {
			return []User{}, errors.Wrapf(err, "http error with url: %s", path)
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
