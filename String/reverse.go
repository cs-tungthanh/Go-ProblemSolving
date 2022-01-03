package main

/*
	Problem: Reverse
*/
// O(n/2)
func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseText(text string, location int) string {
	if location >= 0 {
		return string(text[location]) + reverseText(text, location-1)
	}
	return ""
}

// recursion
func reverseString2(s []byte) {
	s = []byte(reverseText(string(s), len(s)-1))
}

func reverseWords(s string) string {
	bs := []byte(s)
	left := 0
	for i, b := range bs {
		if b == ' ' {
			reverse(&bs, left, i-1)
			left = i + 1
		}
	}
	reverse(&bs, left, len(bs)-1)
	return string(bs)
}

func reverse(src *[]byte, from int, to int) {
	for from < to {
		(*src)[from], (*src)[to] = (*src)[to], (*src)[from]
		from++
		to--
	}
}

// func main() {
// 	arr := []byte("hello")
// 	fmt.Println("Original: ", string(arr))

// 	reverseString2(arr)
// 	fmt.Println(string(arr))

// 	reverseString(arr)
// 	fmt.Println(string(arr))

// 	s := "Let's take LeetCode contest"
// 	fmt.Println(reverseWords(s))
// }
