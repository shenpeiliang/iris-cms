package util

import (
	"regexp"
	"strings"
)

type String struct {
}

//去除包含html的字符串
func (s String) StripTags(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile(`\<[\S\s]+?\>`)
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	//去除STYLE
	re, _ = regexp.Compile(`\<style[\S\s]+?\</style\>`)
	src = re.ReplaceAllString(src, "")

	//去除SCRIPT
	re, _ = regexp.Compile(`\<script[\S\s]+?\</script\>`)
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码
	re, _ = regexp.Compile(`\<[\S\s]+?\>`)
	src = re.ReplaceAllString(src, "")

	//去除空格
	re, _ = regexp.Compile(`\s{1,}`)
	src = re.ReplaceAllString(src, "")

	return src
}

//截取html纯内容
func (s String) SubHtmlText(src string, start int, length int) (substr string) {
	return s.SubStrRune(s.StripTags(src), start, length)
}

//截取字符串
func (s String) SubStrRune(str string, start int, length ...int) (substr string) {
	var (
		runes       = []rune(str)
		runesLength = len(runes)
	)

	if start < 0 {
		start = 0
	}
	if start >= runesLength {
		start = runesLength
	}

	//最大索引
	end := runesLength

	if len(length) > 0 {
		end = start + length[0]
		if end < start {
			end = runesLength
		}
	}
	if end > runesLength {
		end = runesLength
	}
	return string(runes[start:end])
}
