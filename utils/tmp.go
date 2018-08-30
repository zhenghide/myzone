package utils

import "regexp"

func RelpaceHtmlChars(str string) string {
	reg := regexp.MustCompile(`/<[^>]+>/g`)
	rel := reg.ReplaceAllString(str, "")
	return rel
}
