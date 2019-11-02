package parser

type ContentType int

const (
	GitHubContent ContentType = iota
)

type Content struct {
	LinkURL   string
	UserNames []string
	ContentType
}
