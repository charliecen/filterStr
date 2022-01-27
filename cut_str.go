package filterStr

import (
	"regexp"
	"strings"
	"unicode"
)

type cutStr struct {}

var Cut = new(cutStr)

func (c *cutStr) CutString(name string, length int) string {
	// 去除前后空格
	name = strings.TrimSpace(name)
	// 请求参数字符串
	var newName = name
	if len(name) > length {
		// 判断字符串是否是中文
		if c.isChineseChar(name) {
			// 字符串转rune数组，并截取到有效位置
			runeName := []rune(name)
			newName = string(runeName[:length])
		} else {
			// 否则直接截取有效位置
			newName = name[:length]
		}
	}
	return newName
}

// 判断是否中文字符
func (c *cutStr) isChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			return true
		}
	}
	return false
}