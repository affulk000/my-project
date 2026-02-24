// count numbers in a given string
// Pattern to recognize:
// string manipulation(split,trim)
// Approach: Use strings.Fields to split the string into words and count the number of words.
package mockinterview

import "strings"

func countWords(s string) int {
	trimmed := strings.TrimSpace(s)
	if trimmed == "" {
		return 0
	}
	return len(strings.Fields(trimmed))
}
