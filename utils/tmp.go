package utils

import "regexp"

func RelpaceHtmlTag(str string) string {
	reg := regexp.MustCompile(`/<[^>]+>/g`)
	rel := reg.ReplaceAllString(str, "")
	return rel
}
