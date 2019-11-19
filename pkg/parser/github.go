package parser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

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

func (GitHub) parseBody(data []byte) (Content, error) {
	var github entity.GitHub
	s := strings.Replace(string(data), "payload=", "", 1)
	body, err := url.QueryUnescape(s)

	if err != nil {
		return Content{}, errors.Wrapf(err, "url.QueryUnescape got error. body is %v", s)
	}

	if err := json.Unmarshal(bytes.NewBufferString(body).Bytes(), &github); err != nil {
		return Content{}, errors.Wrapf(err, "github json decode error %s", body)
	}
	switch {
	case github.Comment != nil:
		content := Content{
			LinkURL:     github.Comment.HTMLURL,
			UserNames:   userNames(github.Comment.Body),
			ContentType: GitHubMentionContent,
		}
		return content, nil
	case github.Action == entity.GitHubActionTypeAssigned:
		content := Content{
			LinkURL:     github.PullRequest.HTMLURL,
			UserNames:   []string{github.PullRequest.Assignee.Login},
			ContentType: GitHubAssignedContent,
		}
		return content, nil
	default:
		panic(fmt.Sprintf("Unexpected github content pattenr of %v", github))
	}
}
