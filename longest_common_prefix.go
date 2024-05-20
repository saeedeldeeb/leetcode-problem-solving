package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	var prev, current rune
	var common string
	for i := 0; i < getShortestLength(strs); i++ {
		prev = rune(strs[0][i])
		for j := 1; j < len(strs); j++ {
			current = rune(strs[j][i])
			if current != prev {
				return common
			}
			prev = current
		}
		common += string(current)
	}
	return common
}

func getShortestLength(strs []string) int {
	var length int = len(strs[0])
	for _, w := range strs {
		if len(w) < length {
			length = len(w)
		}
	}
	return length
}
