//go:generate mockgen -source=$GOFILE -destination=$GOPATH/src/github.com/bannzai/notifier/pkg/driver/sender_mock_test.go -package=$GOPACKAGE
package driver

import (
	"github.com/bannzai/notifier/pkg/parser"
)

type Sender interface {
	Send(content parser.Content) error
}
