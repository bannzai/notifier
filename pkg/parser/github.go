package parser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bannzai/notifier/pkg/parser/entity"
	"github.com/pkg/errors"
)

type GitHub struct{}

func NewGitHub() GitHub {
	return GitHub{}
}

func (parser GitHub) Parse(request *http.Request) (Content, error) {
	buffer := bytes.Buffer{}
	_, err := buffer.ReadFrom(request.Body)
	if err != nil {
		return Content{}, fmt.Errorf("Request body read error %w", err)
	}

	body := buffer.Bytes()
	return parser.parseBody(body)
}

func (GitHub) parseBody(body []byte) (Content, error) {
	var github entity.GitHub
	if err := json.Unmarshal(body, &github); err != nil {
		return Content{}, errors.Wrapf(err, "github json decode error %s", string(body))
	}

	content := Content{
		LinkURL:     github.Comment.HTMLURL,
		UserNames:   userNames(github.Comment.Body),
		ContentType: GitHubContent,
	}
	return content, nil
}
