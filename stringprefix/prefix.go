package stringprefix

import (
	"strings"
)

func longestCommonPrefix(s []string) string {
	if len(s) < 1 {
		return ""
	}
	prefix := s[0]
	for _, v := range s {
		for strings.Index(v, prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}
