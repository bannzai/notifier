package driver

import (
	"bytes"
	"fmt"
	"net/http"
)

type GitHub struct {
	parameterExtractor
}

func (driver GitHub) Key() string {
	return "github"
}

func (driver GitHub) Drive(r *http.Request) error {
	body := bytes.Buffer{}
	_, err := body.ReadFrom(r.Body)
	if err != nil {
		return fmt.Errorf("Request body read error %w", err)
	}
	driver.parameterExtractor.extract(body.Bytes())
	return nil
}

func NewGitHub() GitHub {
	// TODO: Input default dependency
	return GitHub{}
}
