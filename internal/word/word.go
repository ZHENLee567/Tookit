package word

/*
	在该文档当中定义单词转化的函数
*/
import (
	"bytes"
	"strings"
	"unicode"
)

// 将单词全部转化为大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// 将单词全部转化为小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// 将下划线单词转化为驼峰单词
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s) // 将所有分开单词的首字母大写
	return strings.Replace(s, " ", "", -1)
}

// 将下划线单词转化为小写驼峰单词
func UnderscoreToLowerCamelCase(s string) string {
	// 先转化为驼峰单词，再将首字母小写
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// 驼峰单词转化为下划线单词
func CamelCaseToUnderScore(s string) string {
	// 处理方法：遇到首字母大写的时候，在前面添加下划线
	// 第一个字母前不加下划线
	var output bytes.Buffer
	for i, r := range s {
		if i == 0 {
			output.WriteRune(unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output.WriteRune('_')
		}
		output.WriteRune(unicode.ToLower(r))
	}
	return output.String()
}
