package main

/*
	compute n!
	T(n)= T(n-1) + 3 (3: check if, multiplication, substraction)
		= T(n-2) + 6
		= T(n-k) + 3k
		Til k=n
		= T(0) + 3n
		= 1 + 3n
		= O(n)
	There is no extra space taken during the recursive call,
	-> space complexity is O(n)
*/
func factorial(n int) int {
	if n < 0 {
		return -1
	} else if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}
