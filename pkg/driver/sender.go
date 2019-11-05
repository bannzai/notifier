package driver

import (
	"github.com/bannzai/notifier/pkg/parser"
)

type Sender interface {
	Send(content parser.Content) error
}
