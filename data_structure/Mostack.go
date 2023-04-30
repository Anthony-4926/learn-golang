package data_structure

type MoStack struct {
}

func getNearLessNoRepeat(nums []int) ([]int, []int) {
	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)
	stack := []int{}
	for i, num := range nums {
		for len(stack) != 0 && num < nums[stack[len(stack)-1]] {
			// 出栈
			topI := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 统计栈顶元素左边比它小的
			if len(stack) == 0 {
				left[topI] = -1
			} else {
				left[topI] = stack[len(stack)-1]
			}
			// 统计栈顶元素右边比它小的
			right[topI] = i
		}

		stack = append(stack, i)
	}

	// 复制上一个for循环的循环体即可
	for len(stack) != 0 {
		// 出栈
		topI := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if len(stack) == 0 {
			left[topI] = -1
		} else {
			left[topI] = stack[len(stack)-1]
		}
		// 修改这，改成-1 ！！！
		right[topI] = n
	}

	return left, right
}

func getNearLess(nums []int) ([]int, []int) {
	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)
	stack := [][]int{}

	for i, num := range nums {
		for len(stack) != 0 && num < nums[stack[len(stack)-1][0]] {
			// 出栈
			topList := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			for _, topI := range topList {
				// 统计栈顶元素左边比它小的
				if len(stack) == 0 {
					left[topI] = -1
				} else {
					down := stack[len(stack)-1]
					left[topI] = down[len(down)-1]
				}
				// 统计栈顶元素右边比它小的
				right[topI] = i
			}
		}
		if len(stack) == 0 || num > nums[stack[len(stack)-1][0]] {
			stack = append(stack, []int{i})
		} else {
			stack[n-1] = append(stack[n-1], i)
		}
	}

	// 复制上一个for循环的循环体即可
	for len(stack) != 0 {
		// 出栈
		topList := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, topI := range topList {
			// 统计栈顶元素左边比它小的
			if len(stack) == 0 {
				left[topI] = -1
			} else {
				down := stack[len(stack)-1]
				left[topI] = down[len(down)-1]
			}
			// 统计栈顶元素右边比它小的
			right[topI] = n
		}
	}
	return left, right
}
