package driver

import (
	"net/http"

	"github.com/pkg/errors"
)

type Driver struct {
	Parser
	Sender
}

func New(parser Parser, sender Sender) Driver {
	return Driver{
		Parser: parser,
		Sender: sender,
	}
}

func (driver Driver) Drive(r *http.Request) error {
	content, err := driver.Parser.Parse(r)
	if err != nil {
		return errors.Wrapf(err, "driver.Parser.Parse is failed. request is %v", r)
	}
	return driver.Sender.Send(content)
}
