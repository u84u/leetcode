// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/ 

// this is the naive approach for this problem
// we are basically brute forcing through the entire array
// this is the most inefficient method to do this
// time complexity is O(n^2)
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		a := nums[i]
		for j := i + 1; j < len(nums); j++ {
			b := nums[j]
			sum := a + b
			if sum == target {
				return []int{i + 1, j + 1}
			}
			if sum > target {
				continue
			}
		}
	}
	return []int{}
}


// this is a two pointer approach
// often times the best approach available
// since we already know that the array is sorted we can use this
// we take the low and high end of indexes
// then we loop until we find our answer
// if the sum of left and right index'ed elements is > target we reduce the right pointer
// else we increment the left pointer by one
// we are always going to have a solution
// time complexity is O(n)
func twoSum(nums []int, target int) []int {
    l := 0
    r := len(nums) - 1
    for {
        sum := nums[l] + nums[r]
        if sum > target {
            r -= 1
        } else if sum < target {
            l += 1
        } else {
            return []int{l + 1, r + 1}
        }
    }
    return []int{}
}
