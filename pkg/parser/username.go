package parser

import (
	"regexp"
)

func userNames(target string) []string {
	regex := regexp.MustCompile(`@(\S+)\s*`)
	matches := regex.FindAllStringSubmatch(target, -1)
	if len(matches) == 0 {
		return []string{}
	}
	usernames := []string{}
	for _, matched := range matches {
		usernames = append(usernames, matched[1])
	}
	return usernames
}
