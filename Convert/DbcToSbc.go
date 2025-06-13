package Convert

import (
	"unicode"
)

// DbcToSbc 半角转全角.
func (*Convert) DbcToSbc(input string) string {
	var result []rune

	for _, r := range input {
		if r == ' ' {
			// 半角空格转换为全角空格
			result = append(result, '　')
		} else if unicode.Is(unicode.Han, r) {
			// 汉字保持不变
			result = append(result, r)
		} else {
			// 其他半角字符转换为全角字符
			result = append(result, r+0xFEE0)
		}
	}

	return string(result)

	//return width.Widen.String(s)
}
