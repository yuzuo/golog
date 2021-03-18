package traceback

import (
	"regexp"
	"strings"
)

func GetLogLevelStr(content string) string {
	reg := regexp.MustCompile(`(?i)(INFO|ERROR|DEBUG|WARN|TRACE|FATAL)`)
	params := reg.FindStringSubmatch(content)
	if len(params) > 0 {
		return strings.ToUpper(params[0])
	}
	return ""
}

func GetLogTime(content string) string {
	reg := regexp.MustCompile(`[0-9]{4}[-|/|]?[0-9]{2}[-|/]?[0-9]{2}[ |:|/]?[0-9]{2}[ |:|/]?[0-9]{2}[ |:|/]?[0-9]{2}`)
	params := reg.FindStringSubmatch(content)
	if len(params) > 0 {
		return params[0]
	}
	return ""
}
