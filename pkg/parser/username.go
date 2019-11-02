package parser

import "regexp"

func userNames(target string) []string {
	regex := regexp.MustCompile(`@(\S+)\s+`)
	matched := regex.FindAllStringSubmatch(target, -1)[0]
	usernames := matched[1 : len(matched)-1]
	return usernames
}
