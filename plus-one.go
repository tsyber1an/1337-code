package main

// https://leetcode.com/problems/plus-one/
func plusOne(digits []int) []int {
	sum := 1
	for j := len(digits) - 1; j >= 0; j-- {
		if digits[j]+sum > 9 {
			digits[j] = 0
			if j == 0 {
				digits = append([]int{1}, digits...)
			}
		} else {
			digits[j] = digits[j] + 1
			break
		}
	}

	return digits
}
