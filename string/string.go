package strutil

import (
	"strings"
	"unicode"
)

//RemoveExceptDigit 去除非数字字符
//remove characters in string except digit
func RemoveExceptDigit(origin string) string {
	b := strings.Builder{}
	for _, v := range origin {
		if unicode.IsDigit(v) {
			b.WriteRune(v)
		}
	}
	return b.String()
}

//RemoveExceptDigit2 去除非数字字符
//implemtation using strings.Map
func RemoveExceptDigit2(origin string) string {
	return strings.Map(func(r rune) rune {
		if !unicode.IsDigit(r) {
			return -1
		}
		return r
	}, origin)
}
