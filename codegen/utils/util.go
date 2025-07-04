package utils

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func GetLen(s string) int {
	if s == "" {
		return 0
	}
	start := strings.Index(s, "(")
	end := strings.Index(s, ")")
	if start == -1 || end == -1 || start >= end {
		return 0
	}
	lenstr := s[start+1 : end]
	len, err := strconv.Atoi(lenstr)
	if err != nil {
		return 0
	}
	return len
}

func CamelCase(s string) string {
	if s == "" {
		return ""
	}
	var result strings.Builder
	var upperNext bool

	for i, r := range s {
		if r == '_' {
			upperNext = true
			continue
		}

		if upperNext {
			result.WriteRune(unicode.ToUpper(r))
			upperNext = false
		} else {
			if i == 0 {
				result.WriteRune(unicode.ToLower(r))
			} else {
				result.WriteRune(r)
			}
		}
	}

	return result.String()
}

func PascalCase(s string) string {
	if s == "" {
		return ""
	}
	parts := strings.Split(s, "_")
	for i, part := range parts {
		if len(part) > 0 {
			// 将每个部分的首字母转换为大写
			runes := []rune(part)
			runes[0] = unicode.ToUpper(runes[0])
			parts[i] = string(runes)
		}
	}

	// 拼接所有部分为一个完整的字符串
	return strings.Join(parts, "")
}

func SnakeCase(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	for i, r := range runes {
		if unicode.IsUpper(r) {
			if i > 0 {
				runes[i] = '_' + unicode.ToLower(r)
			} else {
				runes[i] = unicode.ToLower(r)
			}
		}
	}
	return string(runes)
}

// 全局编译正则表达式（避免重复编译）
var (
	camelCaseRegex         = regexp.MustCompile(`([a-z0-9])([A-Z])`)
	nonAlphaNumericRegex   = regexp.MustCompile(`[^a-zA-Z0-9]+`)
	consecutiveDashesRegex = regexp.MustCompile(`-+`)
)

func KebabCase(s string) string {
	s = camelCaseRegex.ReplaceAllString(s, "${1}-${2}")

	s = nonAlphaNumericRegex.ReplaceAllString(s, "-")

	s = strings.ToLower(s)

	s = consecutiveDashesRegex.ReplaceAllString(s, "-")

	// 步骤5: 移除首尾的连字符
	s = strings.Trim(s, "-")

	return s
}
