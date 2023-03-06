package data_structure

import "fmt"

func RadisSort(nums []int, L int) []int {

	count := make([]int, 10)
	help := make([]int, len(nums))

	for i := 1; i <= L; i++ {

		for _, v := range nums {
			count[getDigit(v, i)]++
		}

		for j := 1; j < len(count); j++ {
			count[j] += count[j-1]
		}

		for j := len(nums) - 1; j >= 0; j-- {
			d := getDigit(nums[j], i)
			count[d]--
			help[count[d]] = nums[j]
			fmt.Println(help)
		}

		nums, help = help, nums
	}
	return nums
}

func getDigit(num, d int) int {
	for i := 0; i < d-1; i++ {
		num /= 10
	}
	return num % 10
}

func lenMax(nums []int) int {
	res := 0
	for _, v := range nums {
		i := 0
		for v != 0 {
			v /= 10
			i++
		}
		if i > res {
			res = i
		}
	}
	return res
}
