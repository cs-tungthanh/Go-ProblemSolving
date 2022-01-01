package main

// https://leetcode.com/problems/permutation-in-string/
// This solution work only in lower case letter
func checkInclusion(s1 string, s2 string) bool {
	// there are 26 characters in lower case from a->z
	s1Arr, s2Arr := [26]int{}, [26]int{}
	for _, v := range s1 {
		s1Arr[v-'a']++
	}
	for i, v := range s2 {
		if i >= len(s1) {
			s2Arr[s2[i-len(s1)]-'a']--
		}
		s2Arr[v-'a']++
		if s1Arr == s2Arr {
			return true
		}
	}
	return false
}

// func main() {
// 	checkInclusion("eidbacooo", "acb") // -> true - there's 'bac' is a permutation of 'abc'
// }
