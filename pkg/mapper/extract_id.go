package mapper

import (
	"fmt"

	"github.com/bannzai/notifier/pkg/parser"
)

func noMatchedError(idMapping IDMapping, content parser.Content, toContentType parser.ContentType) error {
	return fmt.Errorf("Not matched id from content of %v, to %d, with mapping values %v", content, toContentType, idMapping)
}

func ExtractID(idMapping IDMapping, content parser.Content, toContentType parser.ContentType) (string, error) {
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
