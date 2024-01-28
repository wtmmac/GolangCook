package main

import (
	"html"
	"strings"
)

// EscapeHTML 将HTML进行转义
func EscapeHTML(input string) string {
	return html.EscapeString(input)
}

// SanitizeHTML 使用白名单过滤HTML标签
func SanitizeHTML(input string, allowedTags []string) string {
	// 使用strings包的Replace函数替换不在白名单中的标签
	for _, tag := range allowedTags {
		openTag := "<" + tag
		closeTag := "</" + tag + ">"
		input = strings.ReplaceAll(input, openTag, "&lt;"+tag)
		input = strings.ReplaceAll(input, closeTag, "&lt;/"+tag)
	}
	return input
}

func main() {
	// 用户输入
	userInput := "<p>This is <b>bold</b> text.</p><script>alert('Hello');</script>"

	// 对HTML进行转义
	escapedInput := EscapeHTML(userInput)
	println("Escaped HTML:", escapedInput)
	println(html.UnescapeString(escapedInput))

	// 白名单过滤HTML标签
	allowedTags := []string{"p", "b", "i", "u"}
	sanitizedInput := SanitizeHTML(userInput, allowedTags)
	println("Sanitized HTML:", sanitizedInput)
}
