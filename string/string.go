package strutil

import (
	"strings"
	"unicode"
)

// 去除非数字字符
// remove characters in string except digit
func RemoveExceptDigit(origin string) string {
	b := strings.Builder{}
	for _, v := range origin {
		if unicode.IsDigit(v) {
			b.WriteRune(v)
		}
	}
	return b.String()
}
