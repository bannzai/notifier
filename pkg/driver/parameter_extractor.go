package driver

type parameter struct {
	prOrIssueURL string
}

type parameterExtractor struct {
}

func (extractor parameterExtractor) extract(body []byte) parameter {
	return parameter{}
}
