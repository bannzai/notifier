package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bannzai/notifier/pkg/parser/entity"
	"github.com/pkg/errors"
)

type GitHub struct{}

func NewGitHub() GitHub {
	return GitHub{}
}

func (parser GitHub) Parse(request *http.Request) (Content, error) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return Content{}, fmt.Errorf("Request body read error %w", err)
	}
	return parser.parseBody(body)
}

func (GitHub) parseBody(body []byte) (Content, error) {
	var github entity.GitHub
	if err := json.Unmarshal(body, &github); err != nil {
		return Content{}, errors.Wrapf(err, "github json decode error %s", body)
	}
	fmt.Printf("Received github structure = %+v\n", github)
	switch {
	case github.Comment != nil:
		// NOTE: Mention on Comment from Issue or Pull-Request
		content := Content{
			LinkURL:     github.Comment.HTMLURL,
			UserNames:   userNames(github.Comment.Body),
			ContentType: GitHubMentionContent,
		}
		return content, nil
	case github.Action == entity.GitHubActionTypeAssigned:
		// NOTE: Assigned from Issue or Pull-Request
		usernames := []string{}
		for _, assignee := range github.PullRequest.Assignees {
			usernames = append(usernames, assignee.Login)
		}
		content := Content{
			LinkURL:     github.PullRequest.HTMLURL,
			UserNames:   usernames,
			ContentType: GitHubAssignedContent,
		}
		return content, nil
	case github.Action == entity.GitHubActionTypeReviewRequested:
		// NOTE: RequestReviewed from Issue or Pull-Request
		usernames := []string{}
		for _, reviewer := range github.PullRequest.RequestReviewers {
			usernames = append(usernames, reviewer.Login)
		}
		fmt.Printf("usernames = %+v\n", usernames)
		if len(usernames) == 0 {
			fmt.Printf("github.RequestedReviwer.Login = %+v\n", github.RequestedReviwer.Login)
			usernames = append(usernames, github.RequestedReviwer.Login)
		}
		content := Content{
			LinkURL:     github.PullRequest.HTMLURL,
			UserNames:   usernames,
			ContentType: GitHubRequestReviewedContent,
		}
		return content, nil
	default:
		panic(fmt.Sprintf("Unexpected github content pattenr of %v", github))
	}
}
