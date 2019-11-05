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

func (content Content) ExtractID(targetContentType ContentType) string {
	m := content.extractIDMap(content.ContentType)
	return m[targetContentType]
}

func (content Content) extractIDMap(fromContentType ContentType) map[ContentType]string {
	return map[ContentType]string{}
}
