package sender

import (
	"github.com/bannzai/notifier/pkg/parser"
)

type Mapper interface {
	MapIDs(content parser.Content, toContentType ContentType) ([]string, error)
}
