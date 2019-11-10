package parser

type ContentType int

const (
	GitHubMentionContent ContentType = iota
	GitHubAssignedContent
	SlackContent
)

type Content struct {
	LinkURL   string
	UserNames []string
	ContentType
}
