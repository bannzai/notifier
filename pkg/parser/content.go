package parser

type ContentType int

const (
	GitHubMentionContent ContentType = iota
	GitHubAssignedContent
	GitHubRequestReviewedContent
)

type Content struct {
	LinkURL   string
	UserNames []string
	ContentType
}
