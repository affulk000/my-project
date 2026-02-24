// This uses the Boyer–Moore bad character rule.
// I compare the pattern from right to left.
// On a mismatch,
// I shift the pattern based on the last occurrence of the mismatched character,
// which reduces unnecessary comparisons.
package mockinterview

func boyerMoore(text, pattern string) []int {
	var result []int

	n, m := len(text), len(pattern)
	if m == 0 || m > n {
		return result
	}

	// 1. Build bad character table
	badChar := make(map[byte]int)
	for i := 0; i < m; i++ {
		badChar[pattern[i]] = i
	}

	// 2. search
	i := 0
	for i <= n-m {
		j := m - 1

		// compare from right to left
		for j >= 0 && text[i+j] == pattern[j] {
			j--
		}

		// match found
		if j < 0 {
			result = append(result, i)
			i += m
		} else {
			last, ok := badChar[text[i+j]]
			if ok {
				i += max(1, j-last)
			} else {
				i += j + 1
			}
		}
	}

	return result
}
