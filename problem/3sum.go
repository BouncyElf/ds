package problem

import "sort"

/*
问题：
15.三数之和
https://leetcode-cn.com/problems/3sum/
*/

// 思路是通过排序+双指针把时间复杂度降低顺便去重
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := range nums {
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, len(nums)-1; j < k; {
			// 去重
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				// 此时i与j都确定，k唯一，所以需要移动jk两个指针
				j++
				k--
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return res
}
