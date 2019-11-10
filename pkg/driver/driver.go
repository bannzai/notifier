package driver

import (
	"net/http"

	"github.com/pkg/errors"
)

type Driver interface {
	Drive(r *http.Request) error
}

type DriverImpl struct {
	Parser
	Sender
}

func New(parser Parser, sender Sender) DriverImpl {
	return DriverImpl{
		Parser: parser,
		Sender: sender,
	}
}

func (driver DriverImpl) Drive(r *http.Request) error {
	content, err := driver.Parser.Parse(r)
	if err != nil {
		return errors.Wrapf(err, "driver.Parser.Parse is failed. request is %v", r)
	}
	return driver.Send(content)
}
