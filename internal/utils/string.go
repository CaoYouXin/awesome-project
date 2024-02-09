package utils

import (
	"unicode/utf8"
)

func SubString(s string, start int, end int) string {
	if start < 0 || end > utf8.RuneCountInString(s) {
		panic("字符串截取参数错误！")
	}

	if start == end {
		return ""
	}

	res := ""

	for i, ch := range []rune(s) {
		if i < start {
			continue
		}

		if i == end {
			break
		}

		res += string(ch)
	}

	return res
}
