package removeDuplicates

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	right := 1
	for left := 1; left < n; left++ {
		if nums[left] != nums[left-1] {
			nums[right] = nums[left]
			right++
		}
	}
	return right
}
