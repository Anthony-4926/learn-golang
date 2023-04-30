package data_structure

import "fmt"

type SegmentTree struct {
	maxN       int
	nums       []int
	sum        []int
	lazy       []int
	updateTo   []int
	isUpdating []bool
}

func NewSegmentTree(orign []int) *SegmentTree {
	t := &SegmentTree{}
	t.maxN = len(orign) + 1
	t.nums = make([]int, t.maxN) // nums[0]不用，从1开始
	copy(t.nums[1:], orign)

	size := t.maxN << 2
	t.sum = make([]int, size)         // 用来支持脑补概念中，某一个范围的累加和信息
	t.lazy = make([]int, size)        // 用来支持脑补概念中，某一个范围没有往下传递的累加任务
	t.updateTo = make([]int, size)    // 用来支持脑补概念中，某一个范围没有更新操作的任务
	t.isUpdating = make([]bool, size) // 用来支持脑补概念中，某一个范围更新任务，更新成了什么

	return t
}

// pushUp 更新根的累加和值
func (t *SegmentTree) pushUp(rt int) {
	t.sum[rt] = t.sum[rt<<1] + t.sum[rt<<1|1]
}

// Build 构建线段树
// 数组中[l, r]范围，对应树中根rt
func (t *SegmentTree) Build(l, r, rt int) {
	if l == r {
		t.sum[l] = t.nums[l]
		return
	}

	mid := (l + r) >> 1
	t.Build(l, mid, rt<<1)
	t.Build(mid+1, r, rt<<1|1)
	t.pushUp(rt)
}

// Add [L, R]范围上，全都加个C，这个任务是永远都不变的，
// 向下传递的时候，也是这个范围
// 当前来到的根是rt，rt包括的范围是[l, r]
func (t *SegmentTree) Add(L, R, C int, l, r int, rt int) {
	// [L, R]范围已经包含了[l, r]
	if L <= l && R >= r {
		t.sum[rt] += C * (r - l + 1) // rt 负责的范围，每一个数都加C，即一共加（r - l + 1）个C
		t.lazy[rt] += C              // 之前懒的，加上这次懒的
		return
	}
	left := rt << 1
	right := left | 1
	mid := (l + r) >> 1
	// 如果任务没有覆盖rt负责的范围，懒不下去了，将懒信息下发一层，然后下发任务
	t.pushDown(rt, mid-l+1, r-mid)
	if L <= mid {
		t.Add(L, R, C, l, mid, left)
	}
	if R > mid {
		t.Add(L, R, C, mid+1, r, right)
	}
	t.pushUp(rt)
}

// Update [L, R]范围上，全都变成C，这个任务是永远都不变的，
// 向下传递的时候，也是这个范围
// 当前来到的根是rt，rt包括的范围是[l, r]
func (t *SegmentTree) Update(L, R, C int, l, r int, rt int) {
	if L <= l && R >= r {
		t.isUpdating[rt] = true
		t.updateTo[rt] = C
		t.sum[rt] = (r - l + 1) * C
		t.lazy[rt] = 0
		return
	}
	left := rt << 1
	right := left | 1
	mid := (l + r) >> 1
	// 如果任务没有覆盖rt负责的范围，懒不下去了，将懒信息下发一层，然后下发任务
	t.pushDown(rt, mid-l+1, r-mid)
	if L <= mid {
		t.Update(L, R, C, l, mid, left)
	}
	if R > mid {
		t.Update(L, R, C, mid+1, r, right)
	}
	t.pushUp(rt)
}

// pushDown 懒不下去时，下发一层
// ln: rt左子树上有几个数
// rn: rt右子树上有几个数
func (t *SegmentTree) pushDown(rt int, ln, rn int) {
	left := rt << 1
	right := left | 1

	if t.isUpdating[rt] {
		t.isUpdating[left] = true
		t.isUpdating[right] = true

		t.updateTo[left] = t.updateTo[rt]
		t.updateTo[right] = t.updateTo[rt]

		t.lazy[left] = 0
		t.lazy[right] = 0

		t.sum[left] = t.updateTo[rt] * ln
		t.sum[right] = t.updateTo[rt] * rn

		t.isUpdating[rt] = false
	}
	if t.lazy[rt] != 0 {
		// 下发左子树，左子树懒
		t.lazy[left] += t.lazy[rt]
		t.sum[left] += ln * t.lazy[rt]
		// 下发右子树，右子树懒
		t.lazy[right] += t.lazy[rt]
		t.sum[right] += rn * t.lazy[rt]

		// 清空自己懒
		t.lazy[rt] = 0
	}
}

func (t *SegmentTree) Query(L, R int, l, r int, rt int) int {
	left := rt << 1
	right := left | 1
	if L <= l && R >= r {
		return t.sum[rt]
	}
	mid := (l + r) >> 1
	t.pushDown(rt, mid-l+1, r-mid)
	ans := 0
	if L <= mid {
		ans += t.Query(L, R, l, mid, left)
	}
	if R > mid {
		ans += t.Query(L, R, mid+1, r, right)
	}
	return ans
}
func main() {
	origin := []int{2, 1, 1, 2, 3, 4, 5}

	//	使用线段树的准备工作
	st := NewSegmentTree(origin)

	// S: 整个区间的开始位置，规定从1开始，不从0开始 -> 固定
	// N: 整个区间的结束位置，规定能到N，不是N-1 -> 固定
	// rt: 整棵树的头节点位置，规定是1，不是0 -> 固定
	S, N, rt := 1, len(origin), 1

	// L: 操作区间的开始位置 -> 可变
	// R:  操作区间的结束位置 -> 可变
	// C: 要加的数字或者要更新的数字 -> 可变
	L, R, C := 2, 5, 4

	// 区间生成，必须在[S,N]整个范围上build
	st.Build(S, N, rt)

	// 区间更新，可以改变L、R和C的值，其他值不可改变
	st.Add(L, R, C, S, N, rt)
	// 区间查询，可以改变L和R的值，其他值不可改变
	sum := st.Query(L, R, S, N, rt)

	fmt.Println(sum)
}
