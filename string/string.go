package strutil

import (
	"strings"
	"unicode"
)

func RemoveExceptDigit(origin string) string {
	b := strings.Builder{}
	for _, v := range origin {
		if unicode.IsDigit(v) {
			b.WriteRune(v)
		}
	}
	return b.String()
}
