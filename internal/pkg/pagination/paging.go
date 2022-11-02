package paging

import "strings"

// SeekSign returns >= for ASC and <= by default.
func SeekSign(sortOrder string) string {
	if strings.ToUpper(sortOrder) == "ASC" {
		return ">="
	}
	return "<="
}
