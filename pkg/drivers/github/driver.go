package github

import (
	"bytes"
	"fmt"
	"net/http"
)

type Driver struct {
	parameterExtractor
}

func (driver Driver) Key() string {
	return "github"
}

func (driver Driver) Drive(r *http.Request) error {
	body := bytes.Buffer{}
	_, err := body.ReadFrom(r.Body)
	if err != nil {
		return fmt.Errorf("Request body read error %w", err)
	}
	driver.parameterExtractor.extract(body.Bytes())
	return nil
}

func NewDriver() Driver {
	// TODO: Input default dependency
	return Driver{}
}
