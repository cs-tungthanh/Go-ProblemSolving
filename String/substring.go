package main

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	left, max_len := 0, 0
	m := make(map[rune]int)
	for k, v := range s {
		if _, ok := m[v]; ok {
			left = max(left, m[v]+1)
		}
		max_len = max(max_len, k-left+1)
		m[v] = k
	}
	return max_len
}
