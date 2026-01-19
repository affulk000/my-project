package exercise

func reverseString(s string) string {
	// convert string to rune slice
	runes := []rune(s)

	// reverse the slice
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
