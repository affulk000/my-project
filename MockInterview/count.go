// count numbers in a given string
package mockinterview

import "strings"

func countWords(s string) int {
	trimmed := strings.TrimSpace(s)
	if trimmed == "" {
		return 0
	}
	return len(strings.Fields(trimmed))
}
