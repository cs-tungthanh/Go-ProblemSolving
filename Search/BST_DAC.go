package main

import "fmt"

/*
	Problem: https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
	Solution:
	- Binary Search (BST)
	- Divide and Conquer (DAC)
*/
func searchRange1(nums []int, target int) []int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			i, d := mid, mid
			for (i+1 <= len(nums)-1) && (nums[i+1] == target) {
				i++
			}
			for (d-1 >= 0) && (nums[d-1] == target) {
				d--
			}
			return []int{d, i}
		}
	}
	return []int{-1, -1}
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	return divideAndConquer(nums, 0, len(nums)-1, target)
}

func divideAndConquer(nums []int, l, r, target int) []int {
	if nums[l] == target && nums[r] == target {
		return []int{l, r}
	}
	if nums[l] <= target && target <= nums[r] {
		mid := l + (r-l)/2
		left := divideAndConquer(nums, l, mid, target)
		right := divideAndConquer(nums, mid+1, r, target)
		if left[0] == -1 {
			return right
		}
		if right[0] == -1 {
			return left
		}
		return []int{left[0], right[1]}
	}
	return []int{-1, -1}
}

func main() {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 6, 6, 6, 8, 10, 10}
	fmt.Println("length", len(arr))
	fmt.Println(searchRange1(arr, 4))
	fmt.Println(searchRange(arr, 4))
}
