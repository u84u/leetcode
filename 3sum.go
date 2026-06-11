// https://leetcode.com/problems/3sum/

// the first medium problem
// this is a dual problem in one
// as in for this, we first sort the array
// so if there are duplicates (there always are) we can skip them
// so after sorting, we loop over the entire array, fixing one element at a time
// then we check if the current element we are at is equal to its predecessor if index > 0, to skip duplicates
// then for each element we create left and right pointer, pointing to element index + 1 and at the end of the array
// then its just a simple Two Sum II problem but the target is 0.
// note that we also have an inner loop that increments the left pointer when finishing until we skip all duplicates
// time complexity is O(n logn) + O(n^2) = O(n^2)
import "sort"
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		a := nums[i]
		if i > 0 && a == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			threeSum := a + nums[l] + nums[r]
			if threeSum > 0 {
				r -= 1
			} else if threeSum < 0 {
				l += 1
			} else {
				res = append(res, []int{a, nums[l], nums[r]})
				l += 1
				for l < r && nums[l] == nums[l-1] {
					l += 1
				}
			}
		}
	}
    return res
}
