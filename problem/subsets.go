package problem

/*
问题：
78.子集
https://leetcode-cn.com/problems/subsets/
*/

func subsets(nums []int) [][]int {
	res := make([][]int, 0, 1<<len(nums))
	for i := range make([]struct{}, 1<<len(nums)) {
		temp := []int{}
		for j := 0; j < i; j++ {
			if (1<<j)&i > 0 {
				temp = append(temp, nums[j])
			}
		}
		res = append(res, temp)
	}
	return res
}
