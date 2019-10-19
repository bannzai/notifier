package github

type parameter struct {
	prOrIssueURL string
}

type parameterExtractor struct {
}

func (extractor parameterExtractor) extract(url string) parameter {
	return extract{}
}
