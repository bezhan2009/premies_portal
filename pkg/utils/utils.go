package utils

import "strings"

func ReplaceBackslashWithSlash(s string) string {
	return strings.ReplaceAll(s, "\\", "/")
}
