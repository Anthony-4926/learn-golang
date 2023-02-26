package sort

func mergeSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) >> 1
	mergeSort(nums, l, mid)
	mergeSort(nums, mid+1, r)
	merge(nums, l, mid, r)
}

func merge(nums []int, l, mid, r int) {
	help := make([]int, 0, r-l+1)
	p, q := l, mid+1

	for p <= mid && q <= r {
		if nums[p] < nums[q] {
			help = append(help, nums[p])
			p++
		} else {
			help = append(help, nums[q])
			q++
		}
	}

	for p <= mid {
		help = append(help, nums[p])
		p++
	}
	for q <= r {
		help = append(help, nums[q])
		q++
	}
	copy(nums[l:], help)
}
