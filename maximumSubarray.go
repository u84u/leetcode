// https://leetcode.com/problems/maximum-subarray

// this is the first sliding window question
// a growing sliding window as the subarray could be anywhere from 1-(n-1) in length
// we are basically looping through the entire array
// and removing the window start ptr
// if the prefix becomes negative
// as negative prefix would only be making the subarray sum smaller
// time complexity is O(n)
func maxSubArray(nums []int) int {
	maxSub := nums[0]
	sum := 0
	for i := 0; i < len(nums); i++ {
		if sum < 0 {
			sum = 0
		}
		sum += nums[i]
		maxSub = max(maxSub, sum)
	}
	return maxSub
}
