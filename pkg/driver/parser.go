package driver

import (
	"net/http"

	"github.com/bannzai/notifier/pkg/parser"
)

type Parser interface {
	Parse(request *http.Request) (parser.Content, error)
}
