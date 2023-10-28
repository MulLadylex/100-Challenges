package removeElement

func removeElement(nums []int, val int) int {
	i, j := 0, 0

	for i < len(nums) {
		if nums[i] == val {
			i++
		} else {
			nums[j] = nums[i]
			i++
			j++
		}
	}

	return j
}

// 官方题解
func removeElement1(nums []int, val int) int {
	left := 0
	for _, v := range nums {
		if v != val {
			nums[left] = v
			left++
		}
	}

	return left
}

// 优化
func removeElement2(nums []int, val int) int {
	left, right := 0, len(nums)
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}

	return left
}
