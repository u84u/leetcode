// https://leetcode.com/problems/contains-duplicate/


// this is a naive approach, we are brute forcing the solution
// this is an inefficient approach to this problem
// time complexity is O(n^2)
func containsDuplicate(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}


// a better approach would be to first
// sort the array and then compare adjacent elements
// since sorting isnt free of time complexity
// the final time complexity would be O(n log(n))
import "slices"
func containsDuplicate(nums []int) bool {
	slices.Sort(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}


// here is the most optimal solution
// however note that, if N is small enough
// like on leetcode submissions, the above approach
// of sorting and checking is superior as it has better cache locallity
// if N is large enough the hashmap approach would be a better choice
// time complexity is O(n)
func containsDuplicate(nums []int) bool {
    hm := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
        num := nums[i]
		_, ok := hm[num]
        if ok { return true }
        hm[num] = struct{}{}
	}
	return false
}
