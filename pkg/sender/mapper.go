package sender

import (
	"github.com/bannzai/notifier/pkg/parser"
)

type Mapper interface {
	MapID(content parser.Content, toContentType parser.ContentType) (string, error)
}
