package mapper

import (
	"fmt"

	"github.com/bannzai/notifier/pkg/parser"
	"github.com/bannzai/notifier/pkg/sender"
)

type Mapper struct{}

func New() Mapper {
	return Mapper{}
}

func noMatchedError(users []User, content parser.Content, toContentType sender.ContentType) error {
	return fmt.Errorf("Not matched id from content of %v, to %d, with users %v", content, toContentType, users)
}

func (Mapper) MapIDs(content parser.Content, toContentType sender.ContentType) ([]string, error) {
	users, err := fetchUsers()
	if err != nil {
		return []string{}, fmt.Errorf("fetchIDMap error %w", err)
	}
	switch content.ContentType {
	case parser.GitHubMentionContent, parser.GitHubAssignedContent:
		ids := []string{}
		for _, username := range content.UserNames {
			slack, ok := extractUserFromGitHub(users, username, toContentType)

			if !ok {
				return ids, noMatchedError(users, content, toContentType)
			}

			if len(slack.ID) > 0 {
				ids = append(ids, slack.ID)
			} else {
				ids = append(ids, slack.Name)
			}
		}
		return ids, nil
	default:
		return []string{}, noMatchedError(users, content, toContentType)
	}
}
