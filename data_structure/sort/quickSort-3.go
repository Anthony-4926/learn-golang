package sort

import (
	"learn-golang/utils"
	"math/rand"
)

func quickSort2(nums []int, L, R int) {
	if L >= R {
		return
	}
	p := partation(nums, L, R)
	quickSort1(nums, L, p-1)
	quickSort1(nums, p+1, R)
}

func partation2(nums []int, L, R int) int {

	randP := rand.Intn(R-L) + L
	utils.Swap(nums, L, randP)
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
