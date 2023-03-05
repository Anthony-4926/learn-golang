package sort

// index: 从index节点向下做heapify
func heapify(nums []int, index, heapSize int) {
	left := index*2 + 1

	for left < heapSize {
		largestIndex := index
		if nums[left] > nums[largestIndex] {
			largestIndex = left
		}
		if left+1 < heapSize && nums[left+1] > nums[largestIndex] {
			largestIndex = left + 1
		}
		if index != largestIndex {
			nums[index], nums[largestIndex] = nums[largestIndex], nums[index]
			index = largestIndex
			left = index*2 + 1
		} else {
			break
		}
	}
}

func heapSort(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		heapify(nums, i, len(nums))
	}
	for i := len(nums) - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, 0, i)
	}
}
