// https://leetcode.com/problems/two-sum

// this is the naive approach, we loop over the array
// and check each element with each other element
// until we find our answer. this is the worst possible way to complete this task.
// time complexity is O(n^2)
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if sum == target {
				return []int{i, j}
			}
		}
	}
	return []int{0, 1}
}


// this is the optimized approach
// we start with an empty hashmap
// and then loop over the array
// for each element we check if target - element is in the hashmap
// if yes we return the correct indicies
// if not we add the element to the hashmap such that
// hashmap[index] = value
// time complexity is O(n)
func twoSum(nums []int, target int) []int {
	hm := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		val := nums[i]
		minus := target - val
		idx, ok := hm[minus]
		if ok {
			return []int{i, idx}
		}
		hm[val] = i
	}
	return []int{0, 1}
}
