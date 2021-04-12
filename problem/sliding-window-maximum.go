package problem

/*
问题：
239.滑动窗口最大值
https://leetcode-cn.com/problems/sliding-window-maximum/
*/

// 维护一个双端单调递减队列，队列中元素为数组下标
func maxSlidingWindow(nums []int, k int) []int {
	q := make([]int, 0, k+1)
	res := make([]int, 0, len(nums)-k+1)
	for i := range nums {
		// 当前元素如果大于等于队尾元素，则pop一个队尾元素
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		// 当前下标放到队尾
		q = append(q, i)

		// 若当前队头元素已经超过滑动窗口左侧，则pop队头元素
		for len(q) > 0 && q[0] <= i-k {
			q = q[1:]
		}
		if i >= k-1 {
			res = append(res, nums[q[0]])
		}
	}
	return res
}
