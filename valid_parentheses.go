package main

func isValid(s string) bool {
	paras := map[rune]rune{'(': ')', '{': '}', '[': ']'}
	var st stack
	var stValue rune
	for _, v := range s {
		if _, ok := paras[v]; ok {
			st = st.Push(v)
			continue
		}
		st, stValue = st.Pop()
		if stValue == rune(0) {
			st = st.Push(v)
			continue
		}
		if paras[stValue] == v {
			continue
		}
		return false
	}
	return len(st) == 0
}

type stack []rune

func (s stack) Push(v rune) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, rune) {
	if len(s) == 0 {
		return s, rune(0)
	}

	l := len(s)
	return s[:l-1], s[l-1]
}
