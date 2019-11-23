//go:generate mockgen -source=$GOFILE -destination=$GOPATH/src/github.com/bannzai/notifier/pkg/driver/parser_mock_test.go -package=$GOPACKAGE
package driver

import (
	"net/http"

	"github.com/bannzai/notifier/pkg/parser"
)

type Parser interface {
	Parse(request *http.Request) (parser.Content, error)
}
