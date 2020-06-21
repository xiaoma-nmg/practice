/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	res := []int{}
	has := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		other := target - nums[i]
		if index, ok := has[other]; ok {
			res = append(res, index, i)
			return res
		}
		has[nums[i]] = i
	}
	return res
}
// @lc code=end

