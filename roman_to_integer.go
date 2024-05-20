package main

func romanToInt(s string) int {
	romanMap := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	count := 0
	for i := 0; i < len(s); {
		if i != len(s)-1 {
			if romanMap[rune(s[i])] < romanMap[rune(s[i+1])] {
				count = count + (romanMap[rune(s[i+1])] - romanMap[rune(s[i])])
				i += 2
				continue
			}
		}
		count = count + romanMap[rune(s[i])]
		i++
	}

	return count
}
