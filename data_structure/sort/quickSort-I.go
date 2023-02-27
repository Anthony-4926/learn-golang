package sort

func quickSort1(nums []int, L, R int) {
	if L >= R {
		return
	}
	p := partation(nums, L, R)
	quickSort1(nums, L, p-1)
	quickSort1(nums, p+1, R)
}

func partation(nums []int, L, R int) int {
	t := nums[L]
	l, r := L, R

	for l < r {
		for r > l && nums[r] >= t {
			r--
		}
		nums[l] = nums[r]

		for l < r && nums[l] <= t {
			l++
		}
		nums[r] = nums[l]
	}
	nums[l] = t
	return l
}
