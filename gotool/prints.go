package gotool

import "strings"

const NEW_LINE_CHAR = "\n"

// PrintWithNewLine prints elements with new line
func PrintWithNewLine(elements []string) string {
	if len(elements) == 0 {
		return ""
	}

	return strings.Join(elements, NEW_LINE_CHAR)
}
