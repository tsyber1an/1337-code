// https://leetcode.com/problems/remove-element/
func removeElement(nums []int, val int) int {
	n := len(nums)
	t := 1
	k := 0

	for i := 0; i < n; i++ {
		if nums[i] == val {
			// pop from back
			for j := n - t; j > i; j-- {
				if nums[j] != val {
					nums[i], nums[j] = nums[j], nums[i]
					t++
					k++
					break
				}
				t++
			}
		} else {
			k++
		}
	}

	return k
}
