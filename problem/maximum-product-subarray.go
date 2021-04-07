package problem

/*
问题：
152.乘积最大子数组
https://leetcode-cn.com/problems/maximum-product-subarray/
*/

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 维护一个当前的最大值以及最小值，不停的去更新这两个值
	x, y, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		tempX, tempY := x, y
		x = max(nums[i], max(nums[i]*tempX, nums[i]*tempY))
		y = min(nums[i], min(nums[i]*tempX, nums[i]*tempY))
		res = max(res, x)
	}
	return res
}
