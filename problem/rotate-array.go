package problem

/*
问题：
189.旋转数组
https://leetcode-cn.com/problems/rotate-array/
*/

// 先全部翻转一边，然后从0到k-1反转一遍，然后从k到末尾再反转一遍
func rotate(nums []int, k int) {
	if k%len(nums) == 0 {
		return
	}
	k = k % len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(data []int, s, e int) {
	if s >= e || e >= len(data) {
		return
	}
	for i, j := s, e; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
