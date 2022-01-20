package main

import "fmt"

/*
	https://leetcode.com/problems/zigzag-conversion/
	Rearrange a string by following a given pattern
	Param:
		- s: string 	= PAYPALHIR
		- numRows: int 	= 3
	The zigzag pattern
	P   A   R
	A P L I
	Y   H
*/
/*
	Approach 1: Sort by row
	Description:
		- we have a corner case when numRows equals 1 it just return the original string
		we need to have n arrays to store this row
		Let's keep an index so we know which is gonna be an array containing each of these array -> r
		- traverse through our string like native idea of zigzag
		The question is when would we go back with the other direction?
		basically when the row is at the top or the bottom, we change direction to update r index.
		we should actually maybe keep a variable to kind of see our direction -> direction
		when we reach at the bottom let's flip it by -> times minus -1 (*-1)
			Step 1 Step 2 Step3
		r	+1		-1		+1
		r	+1		-1		+1
		r	+1		-1		+1
		PAR
		APLI
		YH
		After this action we have filled these arrays completely
		and then we just need to join this array into string
	Time O(n + k < n + n = 2n) = O(n)
	Space O(n)
	n is a number of characters in string
	k is a number of numRows
	- but k is always smaller than n
*/
func convert1(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	arr := make([][]byte, numRows)
	r, direction := 0, -1
	for i := range s {
		arr[r] = append(arr[r], s[i])
		if r == numRows-1 || r == 0 {
			direction *= -1
		}
		r += direction
	}
	result := ""
	for _, rowValue := range arr {
		result += string(rowValue)
	}
	return result
}

/*
	2nd Approach:
	we're gonna find a pattern to sort it.
	Param:
		- s: string 	= PAYPALHIRINGA -> len = 12
		- numRows: int 	= 4
	The zigzag pattern
	row
	1: P     H      A
	2: A   L I    G
	3: Y A   R  N
	4: P	 I
	row 0: 0 	 6       12		-> 0 +6 +6
	row 1: 1   5 7    11		-> 1 +4 +2 +4
	row 2: 2 4   8 10		    -> 2 +2 +4 +2
	row 3: 3     9			    -> 3 +6
	What is the cycle size? you can see the top and bottom the cycle size is 6
	Cycle:
	 numRows -> cycle
	 1 -> 1 special case
	 2 -> 2
	 3 -> 4
	 4 -> 6
	so how do we get it? you can see it equal to traverse 2 times numRows and ignore the top and bot so it would be 2*numRows - 2
	then cycle = 2*numRows - 2

	Take row 1 for example: row=1 -> 1 5 7 11
	row 1 iterate with cycle -> j = 1 7
	so we have found value at 1, 7
	then how do we find k value at index of 5 and 11 in row 1:
	when j = 1 we need to go forward a cycle and go back 2 times of row index
	for j = 1 go a cycle -> 7 -> go back 2 times of row -> 5
	-> k = j + cycle - 2*row_index
	j = 1 -> k = 1+6 - 2*1 = 5
	j = 7 -> k = 7+6 - 2*1 = 11
	So we have found this pattern

	Time: O(n)
	Space: O(n)
*/

func convert2(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	result := ""
	cycle := 2*numRows - 2
	for i := 0; i < numRows; i++ {
		for j := i; j < len(s); j += cycle {
			result += string(s[j])
			k := j + cycle - 2*i
			if i != 0 && i != numRows-1 && k < len(s) {
				// In this condition we just care about number in diagonal line
				result += string(s[k])
			}
		}
	}
	return result
}

func main() {
	fmt.Println(convert1("PAYPALHIRINGA", 4))
	fmt.Println(convert2("PAYPALHIRINGA", 4))
}
