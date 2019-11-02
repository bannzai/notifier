package parser

import "net/http"

type GitHub struct {
}

func NewGitHub() GitHub {
	return GitHub{}
}

func (GitHub) Parse(request *http.Request) (Content, error) {

}

func (GitHub) parseBody(body []byte) Content {

}
