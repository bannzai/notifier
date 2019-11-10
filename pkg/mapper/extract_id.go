package mapper

import (
	"fmt"

	"github.com/bannzai/notifier/pkg/parser"
)

type Mapper struct{}

func noMatchedError(idMapping IDMapping, content parser.Content, toContentType parser.ContentType) error {
	return fmt.Errorf("Not matched id from content of %v, to %d, with mapping values %v", content, toContentType, idMapping)
}

func (Mapper) MapID(content parser.Content, toContentType parser.ContentType) (string, error) {
	idMapping, err := fetchIDMap()
	if err != nil {
		return "", fmt.Errorf("fetchIDMap error %w", err)
	}
	switch content.ContentType {
	case parser.GitHubMentionContent, parser.GitHubAssignedContent:
		id := idMapping.extractFromGitHub(toContentType)
		if len(id) == 0 {
			return "", noMatchedError(idMapping, content, toContentType)
		}
		return id, nil
	default:
		return "", noMatchedError(idMapping, content, toContentType)
	}
}
