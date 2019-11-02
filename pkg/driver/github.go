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
	buffer := bytes.Buffer{}
	_, err := buffer.ReadFrom(r.Body)
	if err != nil {
		return fmt.Errorf("Request body read error %w", err)
	}

	body := buffer.Bytes()

	fmt.Printf("body: %+v\n", string(body))
	driver.parameterExtractor.extract(body)
	return nil
}

func NewGitHub() GitHub {
	// TODO: Input default dependency
	return GitHub{}
}
