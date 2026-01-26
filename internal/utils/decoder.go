package utils

import (
	"strings"
)

func CleanLogLine(line string) string {

	if strings.Contains(line, "\x00") {
		line = strings.ReplaceAll(line, "\x00", "")
	}

	line = strings.TrimSpace(line)
	line = strings.TrimPrefix(line, "\ufeff")

	return line
}
