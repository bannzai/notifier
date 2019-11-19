package mapper

import (
	"github.com/bannzai/notifier/pkg/sender"
)

func extractUserFromGitHub(
	users []User,
	githubUserName string,
	extractedContentType sender.ContentType,
) (Slack, bool) {
	switch extractedContentType {
	case sender.SlackContentType:
		for _, user := range users {
			if user.GitHub.Login == githubUserName {
				return user.Slack, true
			}
		}
		return Slack{}, false
	default:
		return Slack{}, false
	}
}
