package utils

func Swap(nums []int, a, b int) {
	t := nums[a]
	nums[a] = nums[b]
	nums[b] = t
}
